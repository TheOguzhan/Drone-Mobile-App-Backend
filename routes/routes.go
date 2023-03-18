package routes

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func HandleFileUploads(c *fiber.Ctx) error {
	// Parse the multipart form:
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	// => *multipart.Form

	// Get all files from "documents" key:
	files := form.File["uploaded-file"]
	// => []*multipart.FileHeader

	var file_name []string
	// Loop through files:
	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		// => "tutorial.pdf" 360641 "application/pdf"

		savable_file_name := fmt.Sprintf("%d", time.Now().Unix()) + uuid.NewString() + file.Filename

		file_name = append(file_name, "localhost:3000/static-files/"+savable_file_name)

		// Save the files to disk:
		err := c.SaveFile(file, fmt.Sprintf("./static-files/%s", savable_file_name))

		// Check for errors
		if err != nil {
			return err
		}

	}
	return c.JSON(fiber.Map{"file urls": file_name})
}
