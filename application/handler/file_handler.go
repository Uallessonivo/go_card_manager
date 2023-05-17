package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	UseCase interfaces.FileUseCaseInterface
}

func (h FileHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := h.UseCase.ValidateFile(file); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response, err := h.UseCase.SaveData(file)
	if err != nil {
		return c.Status(500).JSON(err.Error())
	}

	if response.FailedCards != nil {
		return c.Status(400).JSON(response)
	}

	return c.Status(200).JSON(fiber.Map{"Message": response.Message})
}

// TOOD
func (h FileHandler) DownloadCardReport(c *fiber.Ctx) error {
	if err := h.UseCase.GenerateCardsReport(); err != nil {
		return err
	}
	return nil
}

// TODO
func (h FileHandler) DownloadEmployeeReport(c *fiber.Ctx) error {
	if err := h.UseCase.GenerateEmployeesReport(); err != nil {
		return err
	}
	return nil
}
