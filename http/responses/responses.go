package responses

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse is the api format response of an error
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail is the api format response of an error detail
type ErrorDetail struct {
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Details string `json:"details"`
}

// GenerateErrorResponse converts a custom error into json response
func GenerateErrorResponse(code int, err error) ([]byte, error) {
	newErrorResp := &ErrorResponse{
		Error: ErrorDetail{
			Status:  code,
			Error:   err.Error(),
			Details: fmt.Sprint(err),
		},
	}

	return json.Marshal(newErrorResp)
}

// GenerateResponse converts a struct response into json response
func GenerateResponse(data *Transaction) ([]byte, error) {
	return json.Marshal(data)
}
