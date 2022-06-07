package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jmsilvadev/atomics-commits-aio/http/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	url := "http://localhost:8080/transactions"
	router := NewRouter()
	router.Get(url)
	localServer := httptest.NewServer(router)
	defer localServer.Close()

	t.Run("Transactions_Without_Body", func(t *testing.T) {
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", url, nil)
		require.NoError(t, err)
		router.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Result().StatusCode)
	})

	t.Run("Transactions", func(t *testing.T) {
		jsonBody := &responses.Transaction{
			Block: "1",
			Hash:  1,
		}
		b, er := json.Marshal(jsonBody)
		require.NoError(t, er)
		bd := bytes.NewBuffer(b)
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", url, bd)
		require.NoError(t, err)
		router.ServeHTTP(rr, req)
		assert.Equal(t, 201, rr.Result().StatusCode)
	})

	t.Run("Transactions_With_Invalid_Body", func(t *testing.T) {
		var jsonBody interface{} = `"lock: "1"`
		b, er := json.Marshal(jsonBody)
		require.NoError(t, er)
		bd := bytes.NewBuffer(b)
		rr := httptest.NewRecorder()
		req, err := http.NewRequest("POST", url, bd)
		require.NoError(t, err)
		router.ServeHTTP(rr, req)
		assert.Equal(t, 400, rr.Result().StatusCode)
	})

}

func TestNotFoundHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
	w := httptest.NewRecorder()
	NotFoundRoute(w, req)
	res := w.Result()
	assert.Equal(t, res.StatusCode, http.StatusNotFound)
}
