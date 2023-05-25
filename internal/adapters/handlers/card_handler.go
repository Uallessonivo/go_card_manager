package handlers

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type CardHandler struct {
	CardService ports.CardService
}

func (cd CardHandler) CreateCard(c *fiber.Ctx) error {
	var card models.CardRequest

	if err := c.BodyParser(&card); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := cd.CardService.CreateCard(&card)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&result)
}

func (cd CardHandler) ListCards(c *fiber.Ctx) error {
	results, err := cd.CardService.ListAllCards()

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&results)
}

func (cd CardHandler) ListCardsByType(c *fiber.Ctx) error {
	param := c.Params("type")

	results, err := cd.CardService.ListAllCardsByType(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&results)
}

func (cd CardHandler) ListCardsByOwner(c *fiber.Ctx) error {
	param := c.Params("owner")

	results, err := cd.CardService.ListAllCardsByOwner(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&results)
}

func (cd CardHandler) DeleteCard(c *fiber.Ctx) error {
	param := c.Params("id")

	err := cd.CardService.DeleteCard(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}
