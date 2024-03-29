package handlers

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	FileService ports.FileService
}

func (h FileHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	cards, err := utils.ExtractDataFromExcelFile(file)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	response, err := h.FileService.SaveData(cards)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if response.FailedCards != nil {
		return c.Status(400).JSON(response)
	}

	return c.Status(200).JSON(fiber.Map{"Result": response.Message})
}

func (h FileHandler) DownloadCardReport(c *fiber.Ctx) error {
	param := c.Query("type")

	buf, err := h.FileService.GenerateCardsReport(param)

	if err != nil {
		return c.Status(500).SendString("Error generating report.")
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=relatorio.csv")

	return c.Send(buf.Bytes())
}

func (h FileHandler) DownloadEmployeeReport(c *fiber.Ctx) error {
	buf, err := h.FileService.GenerateEmployeesReport()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", "attachment; filename=relatorio.csv")

	return c.Send(buf.Bytes())
}
