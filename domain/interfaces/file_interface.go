package interfaces

import "mime/multipart"

// TODO

type FileUseCaseInterface interface {
	ValidateFile(file *multipart.FileHeader) error
	SaveData(file *multipart.FileHeader) error
	GenerateCardsReport() error
	GenerateEmployeesReport() error
}
