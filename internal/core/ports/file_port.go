package ports

import (
	"bytes"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type FileService interface {
	SaveData(input []*models.CardRequest) (*models.UploadResponse, error)
	GenerateCardsReport(cardType string) (*bytes.Buffer, error)
	GenerateEmployeesReport() (*bytes.Buffer, error)
}
