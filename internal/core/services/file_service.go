package services

import (
	"bytes"
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	ports2 "github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"os"

	"github.com/Uallessonivo/go_card_manager/application/utils"
)

type FileUseCase struct {
	EmployeeRepository ports2.EmployeeRepository
	CardRepository     ports2.CardRepository
	CardService        ports2.CardService
}

func NewFileService(
	e ports2.EmployeeRepository,
	c ports2.CardRepository,
	u ports2.CardService) ports2.FileService {
	return &FileUseCase{
		EmployeeRepository: e,
		CardRepository:     c,
		CardService:        u,
	}
}

func (f FileUseCase) SaveData(input []*entities2.CardRequest) (*entities2.UploadResponse, error) {
	var failedCards []*entities2.CardRequest

	for _, card := range input {
		if _, err := f.CardService.CreateCard(card); err != nil {
			failedCards = append(failedCards, card)
		}
	}

	if len(failedCards) > 0 {
		return &entities2.UploadResponse{
			Message:     "Some cards are not inserted into the postgres",
			FailedCards: failedCards,
		}, nil
	}

	return &entities2.UploadResponse{
		Message: "All cards have been saved in the postgres",
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
