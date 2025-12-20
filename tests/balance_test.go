package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"flip/internal/app"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type UploadCSVResponse struct {
	Status string `json:"status"`
	Data   struct {
		UploadID string `json:"upload_id"`
	} `json:"data"`
}

type BalanceResponse struct {
	Status string `json:"status"`
	Data   struct {
		Balance int `json:"balance"`
	} `json:"data"`
}

func TestBalanceEndpoint(t *testing.T) {
	app, uc := app.New(context.Background(), 1, 0)

	file, err := os.Open("./../noedit_transaction_for_test.csv")
	assert.NoError(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "noedit_transaction_for_test.csv")
	assert.NoError(t, err)

	_, err = io.Copy(part, file)
	assert.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/statements", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	var statementResponse UploadCSVResponse
	err = json.NewDecoder(resp.Body).Decode(&statementResponse)

	assert.NoError(t, err)
	assert.NotEmpty(t, statementResponse.Data.UploadID)

	uploadID := statementResponse.Data.UploadID

	time.Sleep(time.Second * 2)

	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/balance?upload_id=%s", uploadID), nil)
	resp, err = app.Test(req)

	assert.NoError(t, err)

	var balanceResponse BalanceResponse
	err = json.NewDecoder(resp.Body).Decode(&balanceResponse)

	assert.NoError(t, err)

	balance := uc.GetBalanceByUploadID(context.Background(), uploadID)
	assert.Equal(t, balance, balanceResponse.Data.Balance)
	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
