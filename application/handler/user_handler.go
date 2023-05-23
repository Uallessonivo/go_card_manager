package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UseCase interfaces.UserUseCaseInterface
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user entities.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := u.UseCase.CreateUser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByID(c *fiber.Ctx) error {
	param := c.Params("id")

	result, err := u.UseCase.GetUserByID(param)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	param := c.Params("email")

	result, err := u.UseCase.GetUserByEmail(param)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) UpdateUser(c *fiber.Ctx) error {
	param := c.Params("id")
	var user entities.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := u.UseCase.UpdateUser(param, &user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) DeleteUser(c *fiber.Ctx) error {
	param := c.Params("id")

	err := u.UseCase.DeleteUser(param)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}
