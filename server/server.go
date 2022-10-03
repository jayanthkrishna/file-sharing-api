package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type test struct {
	Hello string `json:"hello"`
	Name  string `json:"name"`
}

func main() {
	app := fiber.New()
	app.Post("/pdf", fiberServerHandler)
	app.Get("/getpdf", getPdf)
	app.Post("/sendpdf", sendpdf)

	fmt.Println("Server running at 8081!!!!")
	err := app.Listen(":8081")

	fmt.Println(err)
}

func fiberServerHandler(c *fiber.Ctx) error {

	formheader, err := c.FormFile("photo")

	if err != nil {
		c.JSON(fiber.Map{
			"Error": fiber.StatusBadRequest,
		})
	}

	err = c.SaveFile(formheader, formheader.Filename)

	if err != nil {
		c.JSON(fiber.Map{
			"error": err,
		})
	}

	f, _ := c.MultipartForm()
	v := map[string]interface{}{}
	for i, j := range f.Value {
		v[i] = j[0]
	}

	var t test
	jstring, _ := json.Marshal(v)
	json.Unmarshal(jstring, &t)
	// json.Unmarshal([]byte(v), &t)

	return c.JSON(fiber.Map{
		"formValue": c.FormValue("data"),
		"f.value":   v,
		"t":         t,
	})
}

func getPdf(c *fiber.Ctx) error {

	return c.SendFile("./testdoc.pdf")

}

func sendpdf(c *fiber.Ctx) error {

	var t test = test{
		Hello: "helloClient",
		Name:  "JayanthClient",
	}

	a, _ := json.Marshal(&t)
	res := map[string]string{}
	json.Unmarshal(a, &res)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	for i, j := range res {
		fw, _ := writer.CreateFormField(i)
		io.Copy(fw, strings.NewReader(j))
	}

	return c.JSON(fiber.Map{
		"body": body.String(),
		"res":  res,
	})
}

// func createImage(w http.ResponseWriter, request *http.Request) {
// 	err := request.ParseMultipartForm(32 << 20) // maxMemory 32MB
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	//Access the photo key - First Approach
// 	file, h, err := request.FormFile("photo")
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	tmpfile, err := os.Create(".././" + h.Filename)
// 	defer tmpfile.Close()
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	_, err = io.Copy(tmpfile, file)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(200)
// 	return
// }
