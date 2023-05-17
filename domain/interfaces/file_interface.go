package interfaces

import (
	"mime/multipart"

	"github.com/Uallessonivo/go_card_manager/domain/model"
)

// TODO

type FileUseCaseInterface interface {
	ValidateFile(file *multipart.FileHeader) error
	SaveData(file *multipart.FileHeader) (*model.UploadResponse, error)
	GenerateCardsReport() error
	GenerateEmployeesReport() error
}
