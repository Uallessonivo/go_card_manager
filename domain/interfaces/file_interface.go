package interfaces

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type FileUseCaseInterface interface {
	SaveData(input []*model.CardRequest) (*model.UploadResponse, error)
	GenerateCardsReport() error
	GenerateEmployeesReport() error
}
