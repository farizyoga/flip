package tests

import (
	"bytes"
	"context"
	"flip/internal/app"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadCSVSuccess(t *testing.T) {
	app, _ := app.New(context.Background(), 1, 0)

	file, err := os.Open("./../transaction.csv")
	assert.NoError(t, err)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "transaction.csv")
	assert.NoError(t, err)

	_, err = io.Copy(part, file)
	assert.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/statements", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestUploadCSVWithoutTheFile(t *testing.T) {
	app, _ := app.New(context.Background(), 1, 0)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/statements", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
}
