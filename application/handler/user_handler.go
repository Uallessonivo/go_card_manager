package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UseCase interfaces.UserUseCaseInterface
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, err := u.UseCase.Create(user.Name, user.Email, user.Password, user.SecretKey)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByID(c *fiber.Ctx) error {
	param := c.Params("id")

	result, err := u.UseCase.GetByID(param)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	param := c.Params("email")

	result, err := u.UseCase.GetByEmail(param)

	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) UpdateUser(c *fiber.Ctx) error {
	param := c.Params("id")
	var user model.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, err := u.UseCase.Update(param, user.Name, user.Email, user.Password)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) DeleteUser(c *fiber.Ctx) error {
	param := c.Params("id")

	err := u.UseCase.Delete(param)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON("OK")
}
