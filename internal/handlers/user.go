package handlers

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/config"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	if result := config.DB.Create(&user); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot create user",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
