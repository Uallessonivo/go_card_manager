package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase model.UserUseCase
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, err := u.useCase.Create(&user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}
