package usecase

import (
	"mime/multipart"
	"path/filepath"

	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
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

func (f FileUseCase) SaveData(file *multipart.FileHeader) (*model.UploadResponse, error) {
	fl, err := excelize.OpenFile(file.Filename)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	cards, err := utils.ExtractDataFromExcelFile(fl)
	if err != nil {
		return nil, err
	}

	var failedCards []*model.CardRequest

	for _, card := range cards {
		if _, err := f.CardeUseCase.CreateCard(card); err != nil {
			failedCards = append(failedCards, card)
		}
	}

	if len(failedCards) > 0 {
		return &model.UploadResponse{
			Message:     "Some cards are not inserted into the database",
			FailedCards: failedCards,
		}, nil
	}

	return &model.UploadResponse{
		Message: "All cards have been saved in the database",
	}, nil
}

func (f FileUseCase) GenerateCardsReport() error {
	//TODO implement me
	panic("implement me")
}

func (f FileUseCase) GenerateEmployeesReport() error {
	//TODO implement me
	panic("implement me")
}
