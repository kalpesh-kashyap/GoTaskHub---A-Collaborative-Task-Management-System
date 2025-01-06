package main

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/config"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/middleware"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnectDB()
	config.MigrateDB()
	routes.UserRouter(app)
	routes.UploadRoutesHandler(app)
	routes.DownloadRoutes(app)

	app.Use(middleware.LoggingMiddleware())
	app.Use(middleware.ErrorHandlerMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to GoTaskHub")
	})

	log.Fatal(app.Listen(":8080"))
}
