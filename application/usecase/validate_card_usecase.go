package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type ValidateCardUseCase struct {
	CardRepository     interfaces.CardRepositoryInterface
	EmployeeRepository interfaces.EmployeeRepositoryInterface
}

func NewValidateCardUseCase(c interfaces.CardRepositoryInterface, e interfaces.EmployeeRepositoryInterface) interfaces.CardValidatorInterface {
	return &ValidateCardUseCase{
		CardRepository:     c,
		EmployeeRepository: e,
	}
}

func (v ValidateCardUseCase) ValidateMaxCards(input string) error {
	cards, err := v.CardRepository.ListByOwner(input)

	if err != nil {
		return err
	}

	if len(cards) >= 2 {
		return errors.MaxNumberOfCards
	}

	return nil
}

func (v ValidateCardUseCase) ValidateOwnerExists(input string) (*model.Employee, error) {
	owner, err := v.EmployeeRepository.Get(input)

	if err != nil {
		return nil, errors.OwnerNotFound
	}

	return owner, nil
}
