package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	useCase interfaces.UserUseCaseInterface
}

func UserRoutes(app *fiber.App, us interfaces.UserUseCaseInterface) {
	httpHandler := &UserHandler{
		useCase: us,
	}
	app.Post("/users", httpHandler.CreateUser)
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, err := u.useCase.Create(user.Name, user.Email, user.Password)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}
