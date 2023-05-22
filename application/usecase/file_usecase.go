package usecase

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"

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

// TODO: Improve
func (f FileUseCase) GenerateCardsReport(cardType string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	header := os.Getenv("CSV_HEADER")
	columns := strings.Split(header, ",")
	writer.Write(columns)

	cards, err := f.CardRepository.ListByTYpe(cardType)
	if err != nil {
		return nil, err
	}

	for _, card := range cards {
		row := []string{card.Serial, card.Owner, "", card.Name}
		writer.Write(row)
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		return nil, err
	}

	return buf, nil
}

func (f FileUseCase) GenerateEmployeesReport() error {
	//TODO implement me
	panic("implement me")
}
