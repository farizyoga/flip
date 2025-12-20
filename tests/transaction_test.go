package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"flip/internal/app"
	"flip/internal/entity"
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

type StatementListResponse struct {
	Status string `json:"status"`
	Data   struct {
		Rows []*entity.Statement `json:"rows"`
	} `json:"data"`
}

func TestTransactionIssueEndpointEmptyData(t *testing.T) {
	app, _ := app.New(context.Background(), 1, 0)

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

	assert.NotEmpty(t, uploadID)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/transactions/issues?upload_id=%s", uploadID), nil)
	resp, err = app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	var statementListResponse StatementListResponse
	err = json.NewDecoder(resp.Body).Decode(&statementListResponse)

	assert.Equal(t, 0, len(statementListResponse.Data.Rows))
}

func TestTransactionIssueEndpointReturnsFailedRecord(t *testing.T) {
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

	time.Sleep(time.Second * 1)

	assert.NotEmpty(t, uploadID)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	uc.UpdateAllStatementToFailed(context.Background(), uploadID)

	req = httptest.NewRequest(http.MethodGet, fmt.Sprintf("/transactions/issues?upload_id=%s", uploadID), nil)
	resp, err = app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	var statementListResponse StatementListResponse
	err = json.NewDecoder(resp.Body).Decode(&statementListResponse)

	assert.Equal(t, len(statementListResponse.Data.Rows), 2)
}
