package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func main() {
	call("http://localhost:8081/pdf", "POST")
}

func call(urlPath, method string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("photo", "test.pdf")
	if err != nil {
		return err
	}
	file, err := os.Open("test.pdf")
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	writer.Close()
	req, err := http.NewRequest(method, urlPath, bytes.NewReader(body.Bytes()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, _ := client.Do(req)
	if rsp.StatusCode != http.StatusOK {
		log.Printf("Request failed with response code: %d", rsp.StatusCode)
	}
	return nil
}
