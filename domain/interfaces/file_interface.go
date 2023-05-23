package interfaces

import (
	"bytes"

	"github.com/Uallessonivo/go_card_manager/domain/entities"
)

type FileUseCaseInterface interface {
	SaveData(input []*entities.CardRequest) (*entities.UploadResponse, error)
	GenerateCardsReport(cardType string) (*bytes.Buffer, error)
	GenerateEmployeesReport() (*bytes.Buffer, error)
}
