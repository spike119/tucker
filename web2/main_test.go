package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadTest(t *testing.T) {
	assert := assert.New(t)
	path := "C:/Users/boss/go/src/web2/uploads/image.jpg"
	file, _ := os.Open(path)
	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err2 := os.Stat(uploadFilePath)
	assert.NoError(err2)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}

	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
