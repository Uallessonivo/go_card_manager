package handlers

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	EmployeeService ports.EmployeeService
}

func (h EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var employee models.EmployeeRequest

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.EmployeeService.CreateEmployee(&employee)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}

func (h EmployeeHandler) ListEmployees(c *fiber.Ctx) error {
	results, err := h.EmployeeService.ListEmployees()

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&results)
}

func (h EmployeeHandler) GetEmployee(c *fiber.Ctx) error {
	param := c.Query("query")

	result, err := h.EmployeeService.GetFiltered(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(&result)
}

func (h EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	param := c.Params("id")

	err := h.EmployeeService.DeleteEmployee(param)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON("OK")
}

func (h EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	param := c.Params("id")

	var employee models.EmployeeRequest

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.EmployeeService.UpdateEmployee(param, &employee)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(200).JSON(result)
}
