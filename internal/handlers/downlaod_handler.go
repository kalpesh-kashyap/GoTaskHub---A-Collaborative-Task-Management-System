package handlers

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

var tasksQueue = make(chan string, 10)
var outputDir = "downloads/"

func init() {
	utils.DownloadPool(tasksQueue, outputDir)
}

func DownlaodHandler(c *fiber.Ctx) error {
	var fileUrls struct {
		URLs []string `json:"urls"`
	}

	if err := c.BodyParser(&fileUrls); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": true, "message": "fail to read url"})
	}

	if len(fileUrls.URLs) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "No URLs provided",
		})
	}

	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err = os.Mkdir(outputDir, 0755)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"message": "Failed to create dir" + err.Error(),
			})
		}
	}

	for _, url := range fileUrls.URLs {
		tasksQueue <- url
	}
	close(tasksQueue)

	return c.JSON(fiber.Map{"message": "Files are in queue to download"})

}
