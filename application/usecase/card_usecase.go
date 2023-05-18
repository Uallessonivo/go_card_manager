package usecase

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type CardUseCase struct {
	CardRepository     interfaces.CardRepositoryInterface
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardValidator      interfaces.CardValidatorInterface
}

func NewCardUseCase(
	u interfaces.CardRepositoryInterface,
	c interfaces.EmployeeRepositoryInterface,
	v interfaces.CardValidatorInterface) interfaces.CardUseCaseInterface {
	return &CardUseCase{
		CardRepository:     u,
		EmployeeRepository: c,
		CardValidator:      v,
	}
}

func (c CardUseCase) ListAllCards() ([]*model.CardResponse, error) {
	items, err := c.CardRepository.List()
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) ListAllCardsByType(input string) ([]*model.CardResponse, error) {
	items, err := c.CardRepository.ListByTYpe(input)
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) ListAllCardsByOwner(input string) ([]*model.CardResponse, error) {
	items, err := c.CardRepository.ListByOwner(input)
	if err != nil {
		return nil, err
	}

	results := utils.CardResponse(items)
	return results, nil
}

func (c CardUseCase) CreateCard(input *model.CardRequest) (*model.CardResponse, error) {
	cardOwner, err := c.CardValidator.ValidateCard(input.Owner)
	if err != nil {
		return nil, err
	}

	newCard, err := model.MakeCard(input, cardOwner.Name)
	if err != nil {
		return nil, err
	}

	er := c.CardRepository.Create(newCard)
	if er != nil {
		return nil, er
	}

	response := model.CardResponse{
		ID:     newCard.ID,
		Type:   newCard.Type,
		Owner:  newCard.Owner,
		Serial: newCard.Serial,
	}
	return &response, nil
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
