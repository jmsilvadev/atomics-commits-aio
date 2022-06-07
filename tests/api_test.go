package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/jmsilvadev/atomics-commits-aio/http/responses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestE2E(t *testing.T) {
	url := "http://localhost:8080/transactions"

	t.Run("Transactions_Without_Body", func(t *testing.T) {
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, url, nil)
		assert.NoError(t, err)
		res, err := client.Do(req)
		require.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("Transactions", func(t *testing.T) {
		jsonBody := &responses.Transaction{
			Block: "1",
			Hash:  1,
		}
		b, er := json.Marshal(jsonBody)
		require.NoError(t, er)
		bd := bytes.NewBuffer(b)
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, url, bd)
		assert.NoError(t, err)
		res, err := client.Do(req)
		require.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("Transactions_With_Invalid_Body", func(t *testing.T) {
		var jsonBody interface{} = `"lock: "1"`
		b, er := json.Marshal(jsonBody)
		require.NoError(t, er)
		bd := bytes.NewBuffer(b)
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodPost, url, bd)
		assert.NoError(t, err)
		res, err := client.Do(req)
		require.NoError(t, err)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

}

func TestNotFoundHandler(t *testing.T) {
	res, err := http.Get("http://localhost:8080/notfound")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)
}
