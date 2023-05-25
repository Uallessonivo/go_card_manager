package services

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	ports "github.com/Uallessonivo/go_card_manager/internal/core/ports"
)

type CardUseCase struct {
	CardRepository     ports.CardRepository
	EmployeeRepository ports.EmployeeRepository
}

func NewCardService(
	u ports.CardRepository,
	c ports.EmployeeRepository) ports.CardService {
	return &CardUseCase{
		CardRepository:     u,
		EmployeeRepository: c,
	}
}

func (c CardUseCase) ListAllCards() ([]*entities2.CardResponse, error) {
	items, err := c.CardRepository.List()
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) ListAllCardsByType(input string) ([]*entities2.CardResponse, error) {
	items, err := c.CardRepository.ListByTYpe(input)
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) ListAllCardsByOwner(input string) ([]*entities2.CardResponse, error) {
	items, err := c.CardRepository.ListByOwner(input)
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) CreateCard(input *entities2.CardRequest) (*entities2.CardResponse, error) {
	cardOwner, err := c.ValidateCard(input.Owner)
	if err != nil {
		return nil, err
	}

	newCard, err := entities2.MakeCard(input, cardOwner.Name)
	if err != nil {
		return nil, err
	}

	er := c.CardRepository.Create(newCard)
	if er != nil {
		return nil, er
	}

	response := entities2.CardResponse{
		ID:     newCard.ID,
		Type:   newCard.Type,
		Owner:  newCard.Owner,
		Serial: newCard.Serial,
	}
	return &response, nil
}

func (c CardUseCase) ValidateCard(input string) (*entities2.Employee, error) {
	owner, err := c.EmployeeRepository.Get(input)
	if err != nil {
		return nil, errors.OwnerNotFound
	}

	cards, err := c.CardRepository.ListByOwner(input)
	if err != nil {
		return nil, err
	}

	if len(cards) >= 2 {
		return nil, errors.MaxNumberOfCards
	}

	return owner, nil
}

func (c CardUseCase) DeleteCard(id string) error {
	_, er := c.CardRepository.GetById(id)
	if er != nil {
		return errors.NotFound
	}

	err := c.CardRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
