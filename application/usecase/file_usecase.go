package usecase

import (
	"mime/multipart"
	"path/filepath"

	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/xuri/excelize/v2"
)

type FileUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardRepository     interfaces.CardRepositoryInterface
	CardeUseCase       interfaces.CardUseCaseInterface
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

func (f FileUseCase) SaveData(file *multipart.FileHeader) error {
	fl, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return err
	}
	defer fl.Close()

	// TODO
	cards, err := utils.ExtractDataFromExcelFile(fl)
	if err != nil {
		return err
	}

	// TOOD: save data in database
	for _, card := range cards {
		if _, err := f.CardeUseCase.CreateCard(card); err != nil {
			return err
		}
	}

	return nil
}

func (f FileUseCase) GenerateCardsReport() error {
	//TODO implement me
	panic("implement me")
}

func (f FileUseCase) GenerateEmployeesReport() error {
	//TODO implement me
	panic("implement me")
}
