package handlers

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/models"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

var uploadQueue = make(chan models.UploadTask, 10)

func init() {
	go utils.StartWorkerPool(uploadQueue)
}

func Uploadhandler(c *fiber.Ctx) error {

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to parse form data",
		})
	}

	files := form.File["file"]
	if files == nil || len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "No files uploaded",
		})
	}

	for _, file := range files {

		src, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"errr":    true,
				"message": "Failed to open file",
			})
		}

		defer src.Close()

		fileContent := make([]byte, file.Size)
		_, err = src.Read(fileContent)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to read file",
			})
		}
		task := models.UploadTask{
			FileName: file.Filename,
			Content:  fileContent,
		}

		uploadQueue <- task
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("%d files queued for upload", len(files)),
	})

}
