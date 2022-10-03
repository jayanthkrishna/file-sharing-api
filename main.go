package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/textproto"
)

var jsonStr = []byte(`{"hello" : "world","name":"jayanth"}`)

type test struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

func main() {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, _ := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {"application/json"}})
	part.Write(jsonStr)

	writer.Close()

	// req, _ := http.NewRequest("POST", "http://1.1.1.1/blabla", body)
	// req.Header.Set("Content-Type", "multipart/mixed; boundary="+writer.Boundary())
	var t test
	json.Unmarshal(jsonStr, &t)

	fmt.Println("t :", t)
	fmt.Println("part :", part)

}
