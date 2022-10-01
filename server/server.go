package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Post("/pdf", fiberServerHandler)

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

	return nil
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
