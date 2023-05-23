package usecase

import (
	"bytes"
	"os"

	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
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

func (f FileUseCase) SaveData(input []*entities.CardRequest) (*entities.UploadResponse, error) {
	var failedCards []*entities.CardRequest

	for _, card := range input {
		if _, err := f.CardeUseCase.CreateCard(card); err != nil {
			failedCards = append(failedCards, card)
		}
	}

	if len(failedCards) > 0 {
		return &entities.UploadResponse{
			Message:     "Some cards are not inserted into the database",
			FailedCards: failedCards,
		}, nil
	}

	return &entities.UploadResponse{
		Message: "All cards have been saved in the database",
	}, nil
}

func (f FileUseCase) GenerateCardsReport(cardType string) (*bytes.Buffer, error) {
	header := os.Getenv("CSV_HEADER")

	cards, err := f.CardRepository.ListByTYpe(cardType)
	if err != nil {
		return nil, err
	}

	rows := make([][]string, len(cards))
	for i, card := range cards {
		rows[i] = []string{card.Serial, card.Owner, "", card.Name}
	}

	buf, err := utils.GenerateCSVFile(header, rows)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (f FileUseCase) GenerateEmployeesReport() (*bytes.Buffer, error) {
	header := os.Getenv("CSV_EMPLOYEE_HEADER")

	employees, err := f.EmployeeRepository.List()
	if err != nil {
		return nil, err
	}

	rows := make([][]string, len(employees))
	for i, employee := range employees {
		rows[i] = []string{employee.ID, employee.Cpf, employee.Name}
	}

	buf, err := utils.GenerateCSVFile(header, rows)

	if err != nil {
		return nil, err
	}

	return buf, nil
}
