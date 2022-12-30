package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type CardUseCase struct {
	CardRepository interfaces.CardRepositoryInterface
}

func NewCardUseCase(u interfaces.CardRepositoryInterface) interfaces.CardUseCaseInterface {
	return &CardUseCase{
		CardRepository: u,
	}
}

func (c CardUseCase) List() ([]model.CardResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardUseCase) ListByType(input string) ([]model.CardResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardUseCase) GetByCpf(input string) (*model.CardResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c CardUseCase) Create(input *model.CardRequest) (*model.CardResponse, error) {
	newCard, err := model.MakeCard(input)
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

func (c CardUseCase) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
