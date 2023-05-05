package usecase

import (
	"mime/multipart"
	"path/filepath"

	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
)

type FileUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardRepository     interfaces.CardRepositoryInterface
}

func NewFileUseCase(e interfaces.EmployeeRepositoryInterface, c interfaces.CardRepositoryInterface) interfaces.FileUseCaseInterface {
	return &FileUseCase{
		EmployeeRepository: e,
		CardRepository:     c,
	}
}

// TODO: more validations
func (f FileUseCase) ValidateFile(file *multipart.FileHeader) error {
	if filepath.Ext(file.Filename) != ".xlsx" {
		return errors.FileExtension
	}

	return nil
}

func (f FileUseCase) SaveData(filePath *multipart.FileHeader) error {
	panic("implement me")
}

func (f FileUseCase) GenerateCardsReport() error {
	//TODO implement me
	panic("implement me")
}

func (f FileUseCase) GenerateEmployeesReport() error {
	//TODO implement me
	panic("implement me")
}
