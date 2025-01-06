package handlers

import (
	"GoTaskHub---A-Collaborative-Task-Management-System/config"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/models"
	"GoTaskHub---A-Collaborative-Task-Management-System/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Can not parse JSON"})
	}
	var user models.User
	result := config.DB.Where("email = ?", input.Email).First(&user)
	if result.Error != nil {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}
	return c.JSON(fiber.Map{"token": token})
}
