package services

import (
	"bytes"
	"os"

	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"

	"github.com/Uallessonivo/go_card_manager/application/utils"
)

type FileUseCase struct {
	EmployeeRepository ports.EmployeeRepository
	CardRepository     ports.CardRepository
	CardService        ports.CardService
}

func NewFileService(
	e ports.EmployeeRepository,
	c ports.CardRepository,
	u ports.CardService) ports.FileService {
	return &FileUseCase{
		EmployeeRepository: e,
		CardRepository:     c,
		CardService:        u,
	}
}

func (f FileUseCase) SaveData(input []*models.CardRequest) (*models.UploadResponse, error) {
	var failedCards []*models.CardRequest

	for _, card := range input {
		if _, err := f.CardService.CreateCard(card); err != nil {
			failedCards = append(failedCards, card)
		}
	}

	if len(failedCards) > 0 {
		return &models.UploadResponse{
			Message:     "Some cards are not inserted into the postgres",
			FailedCards: failedCards,
		}, nil
	}

	return &models.UploadResponse{
		Message: "All cards have been saved in the postgres",
	}, nil
}

func (f FileUseCase) GenerateCardsReport(cardType string) (*bytes.Buffer, error) {
	header := os.Getenv("CSV_HEADER")

	cards, err := f.CardRepository.ListByType(cardType)
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
