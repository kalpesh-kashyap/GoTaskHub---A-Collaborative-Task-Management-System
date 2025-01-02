package routes

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/", handlers.CreateUser)
}
