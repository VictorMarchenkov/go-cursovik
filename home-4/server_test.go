package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetHandler(t *testing.T) {
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("error opening file testfile.txt ", err)
	}
	defer func(f http.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("error closing file: testfile.txt", err)
			//return err
		}
	}(file)
	// действия, необходимые для того, чтобы засунуть файл в запрос
	// в качестве мультипарт-формы
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		fmt.Println("error creating new form-data header with the provided field name and file name ", err)
	}
	io.Copy(part, file)
	writer.Close()
	// опять создаем запрос, теперь уже на /upload эндпоинт
	req, _ := http.NewRequest(http.MethodPost, "/files", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	// создаем ResponseRecorder
	rr := httptest.NewRecorder()
	// создаем заглушку файлового сервера. Для прохождения тестов
	// нам достаточно чтобы он возвращал 200 статус
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok!")
	}))
	defer ts.Close()
	uploadHandler := &UploadHandler{
		UploadDir: "files",
		// таким образом мы подменим адрес файлового сервера
		// и вместо реального, хэндлер будет стучаться на заглушку
		// которая всегда будет возвращать 200 статус, что нам и нужна
		HostAddr: ts.URL,
	}
	// опять же, вызываем ServeHTTP у тестируемого обработчика
	uploadHandler.ServeHTTP(rr, req)
	// Проверяем статус-код ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `testfile.txt`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
