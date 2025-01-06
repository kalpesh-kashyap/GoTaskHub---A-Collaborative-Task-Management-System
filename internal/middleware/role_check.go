package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RoleMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Retrieve token from context set by jwtware
		userToken := c.Locals("user")
		if userToken == nil {
			log.Println("JWT token missing in context")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Unauthorized access",
			})
		}

		// Cast token to *jwt.Token from the correct package
		token, ok := userToken.(*jwt.Token)
		if !ok {
			log.Println("Failed to cast token to *jwt.Token")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid token",
			})
		}

		// Extract claims and check the role
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Invalid JWT claims")
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":   true,
				"message": "Invalid token claims",
			})
		}

		log.Println("JWT Claims:", claims)

		// Verify role
		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error":   true,
				"message": "Access denied",
			})
		}

		return c.Next()
	}
}
