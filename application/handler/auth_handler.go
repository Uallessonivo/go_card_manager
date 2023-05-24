package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUseCase interfaces.AuthUseCaseInterface
}

func (ah AuthHandler) Authenticate(c *fiber.Ctx) error {
	var login entities.LoginRequest

	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := ah.AuthUseCase.Login(&login)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&response)
}
