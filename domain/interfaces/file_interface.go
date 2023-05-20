package interfaces

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
)

type FileUseCaseInterface interface {
	SaveData(input []*entities.CardRequest) (*entities.UploadResponse, error)
	GenerateCardsReport() error
	GenerateEmployeesReport() error
}
