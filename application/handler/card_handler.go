package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type CardHandler struct {
	UseCase interfaces.CardUseCaseInterface
}

func (cd CardHandler) CreateCard(c *fiber.Ctx) error {
	var card entities.CardRequest

	if err := c.BodyParser(&card); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := cd.UseCase.CreateCard(&card)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&result)
}

func (cd CardHandler) ListCards(c *fiber.Ctx) error {
	results, err := cd.UseCase.ListAllCards()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(&results)
}

func (cd CardHandler) ListCardsByType(c *fiber.Ctx) error {
	param := c.Params("type")
	results, err := cd.UseCase.ListAllCardsByType(param)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(&results)
}

func (cd CardHandler) ListCardsByOwner(c *fiber.Ctx) error {
	param := c.Params("owner")
	results, err := cd.UseCase.ListAllCardsByOwner(param)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(&results)
}

func (cd CardHandler) DeleteCard(c *fiber.Ctx) error {
	param := c.Params("id")
	err := cd.UseCase.DeleteCard(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}
