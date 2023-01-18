package handler

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type FileHandler struct {
	UseCase interfaces.FileUseCaseInterface
}

// TODO

func (h FileHandler) UploadFile(c *fiber.Ctx) error {
	if err := h.UseCase.SaveData(); err != nil {
		return err
	}
	return nil
}

func (h FileHandler) DownloadCardReport(c *fiber.Ctx) error {
	if err := h.UseCase.GenerateCardsReport(); err != nil {
		return err
	}
	return nil
}

func (h FileHandler) DownloadEmployeeReport(c *fiber.Ctx) error {
	if err := h.UseCase.GenerateEmployeesReport(); err != nil {
		return err
	}
	return nil
}
