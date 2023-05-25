package handlers

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService ports.UserService
}

func (u UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := u.UserService.CreateUser(&user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByID(c *fiber.Ctx) error {
	param := c.Params("id")

	result, err := u.UserService.GetUserByID(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) GetUserByEmail(c *fiber.Ctx) error {
	param := c.Params("email")

	result, err := u.UserService.GetUserByEmail(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) UpdateUser(c *fiber.Ctx) error {
	param := c.Params("id")

	var user models.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := u.UserService.UpdateUser(param, &user)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (u UserHandler) DeleteUser(c *fiber.Ctx) error {
	param := c.Params("id")

	err := u.UserService.DeleteUser(param)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}
