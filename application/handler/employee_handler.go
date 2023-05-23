package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	UseCase interfaces.EmployeeUseCaseInterface
}

func (h EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var employee entities.EmployeeRequest

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.UseCase.CreateEmployee(&employee)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (h EmployeeHandler) ListEmployees(c *fiber.Ctx) error {
	results, err := h.UseCase.ListEmployees()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(&results)
}

func (h EmployeeHandler) GetEmployee(c *fiber.Ctx) error {
	param := c.Query("query")

	result, err := h.UseCase.GetFiltered(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&result)
}

func (h EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	param := c.Params("id")
	err := h.UseCase.DeleteEmployee(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}

func (h EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	param := c.Params("id")
	var employee entities.EmployeeRequest

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.UseCase.UpdateEmployee(param, &employee)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}
