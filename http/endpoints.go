package http

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmsilvadev/atomics-commits-aio/http/responses"
	"github.com/jmsilvadev/atomics-commits-aio/internal/database"
	"github.com/jmsilvadev/atomics-commits-aio/internal/services"
)

const (
	// ContentType is a type of header
	ContentType = "Content-Type"
	// TypeJSON is a type of a content type header
	TypeJSON    = "application/vnd.api+json"
	invalidBody = "invalid body"
)

// NewRouter has all server routes
func NewRouter() *mux.Router {
	ctx := context.Background()
	router := mux.NewRouter().StrictSlash(true)
	log.Println("creating default route")
	router.HandleFunc("/transactions", CreateBlock(ctx)).Methods(http.MethodPost)

	return router
}

// NotFoundRoute is the default response
func NotFoundRoute(w http.ResponseWriter, r *http.Request) {
	err := errors.New("not found")
	jsonResponse, _ := responses.GenerateErrorResponse(404, err)
	w.Header().Set(ContentType, TypeJSON)
	w.WriteHeader(404)
	w.Write(jsonResponse)
}

func CreateBlock(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		itx := &responses.Transaction{}
		if r.Body == nil {
			jsonResponse, _ := responses.GenerateErrorResponse(402, errors.New(invalidBody))
			w.Header().Set(ContentType, TypeJSON)
			w.WriteHeader(400)
			w.Write(jsonResponse)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&itx)
		if err != nil {
			jsonResponse, _ := responses.GenerateErrorResponse(401, errors.New(invalidBody))
			w.Header().Set(ContentType, TypeJSON)
			w.WriteHeader(400)
			w.Write(jsonResponse)
			return
		}

		db := database.NewPostgres(ctx)
		s := services.New(ctx, db)
		data, err := s.CreateTransaction(ctx, itx.ToEntity())
		if err != nil {
			jsonResponse, _ := responses.GenerateErrorResponse(422, err)
			w.Header().Set(ContentType, TypeJSON)
			w.WriteHeader(422)
			w.Write(jsonResponse)
			return
		}

		jsonResponse, err := responses.GenerateResponse(itx.FromEntity(data))
		if err != nil {
			jsonResponse, _ := responses.GenerateErrorResponse(500, err)
			w.Header().Set(ContentType, TypeJSON)
			w.WriteHeader(500)
			w.Write(jsonResponse)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(jsonResponse)
	}
}
