package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type CardHandler struct {
	UseCase interfaces.CardUseCaseInterface
}

func (cd CardHandler) CreateCard(c *fiber.Ctx) error {
	var card model.CardRequest

	if err := c.BodyParser(&card); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, er := cd.UseCase.CreateCard(&card)
	if er != nil {
		return c.Status(400).JSON(er.Error())
	}

	return c.Status(200).JSON(&result)
}

func (cd CardHandler) ListCards(c *fiber.Ctx) error {
	results, err := cd.UseCase.ListAllCards()
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(&results)
}

func (cd CardHandler) DeleteCard(c *fiber.Ctx) error {
	param := c.Params("id")
	err := cd.UseCase.DeleteCard(param)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("OK")
}
