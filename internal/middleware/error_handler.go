package middleware

import "github.com/gofiber/fiber/v2"

func ErrorHandlerMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": true, "message": err.Error()})
		}
		return nil
	}
}
