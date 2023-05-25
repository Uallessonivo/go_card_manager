package handlers

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService ports.AuthService
}

func (ah AuthHandler) Authenticate(c *fiber.Ctx) error {
	var login models.LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := ah.AuthService.Login(&login)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&response)
}
