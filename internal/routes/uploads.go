package routes

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UploadRoutesHandler(app *fiber.App) {
	upload := app.Group("/uploads")
	upload.Post("/", handlers.Uploadhandler)

}
