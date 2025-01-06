package routes

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/handlers"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/middleware"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func UserRouter(app *fiber.App) {
	user := app.Group("/user")
	user.Post("/", handlers.CreateUser)
	user.Post("/login", handlers.Login)

	user.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("a2j3jKLOenfI32JfnkleoIej23nfMdnfWEkj3nfdlQ"),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Unauthorized access",
			})
		},
	}))

	user.Get("/admin", middleware.RoleMiddleware("admin"), func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome admin"})
	})
}
