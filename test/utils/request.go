package utils

import (
	"io"
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http/httptest"
	"net/http"
	"os"
	"path/filepath"
)

func CreateRequest(method string, url string, headers map[string]string, object interface{}) *http.Request {

	body, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
	}

	req := httptest.NewRequest(method, url, bytes.NewReader(body))
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req
}

func UploadFile(url string, headers map[string]string, filePaths []string) *http.Request {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		part, err := writer.CreateFormFile("file", filepath.Base(filePath))
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			panic(err)
		}
	}
	err := writer.Close()
	if err != nil {
		panic(err)
	}
	req := httptest.NewRequest("POST", url, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	return req
}

func UnmarshallResponse(recorder *httptest.ResponseRecorder, object interface{}) error {
	responseBody, err := io.ReadAll(recorder.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, &object)
	if err != nil {
		return err
	}

	return nil
}
