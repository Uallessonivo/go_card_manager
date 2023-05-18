package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type FileUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardRepository     interfaces.CardRepositoryInterface
	CardeUseCase       interfaces.CardUseCaseInterface
}

func NewFileUseCase(
	e interfaces.EmployeeRepositoryInterface,
	c interfaces.CardRepositoryInterface,
	u interfaces.CardUseCaseInterface) interfaces.FileUseCaseInterface {
	return &FileUseCase{
		EmployeeRepository: e,
		CardRepository:     c,
		CardeUseCase:       u,
	}
}

func (f FileUseCase) SaveData(input []*model.CardRequest) (*model.UploadResponse, error) {
	var failedCards []*model.CardRequest

	for _, card := range input {
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
