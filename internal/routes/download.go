package routes

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func DownloadRoutes(app *fiber.App) {
	download := app.Group("/download")
	download.Post("/", handlers.DownlaodHandler)
}
