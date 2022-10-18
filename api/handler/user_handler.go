package handler

import (
	"github.com/Uallessonivo/go_card_manager/api/usecase"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUseCase usecase.UserUseCase
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, err := u.UserUseCase.Create(user.Name, user.Email, user.Password)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}
