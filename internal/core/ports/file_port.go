package ports

import (
	"bytes"
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type FileService interface {
	SaveData(input []*entities2.CardRequest) (*entities2.UploadResponse, error)
	GenerateCardsReport(cardType string) (*bytes.Buffer, error)
	GenerateEmployeesReport() (*bytes.Buffer, error)
}
