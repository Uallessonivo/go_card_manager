package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"github.com/gofiber/fiber/v2"
)

type EmployeeHandler struct {
	UseCase interfaces.EmployeeUseCaseInterface
}

func (h EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var employee model.EmployeeRequest

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	result, er := h.UseCase.CreateEmployee(&employee)
	if er != nil {
		return c.Status(400).JSON(er.Error())
	}

	return c.Status(200).JSON(result)
}
