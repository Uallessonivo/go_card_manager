package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type CardRepositoryInterface interface {
	List() ([]model.Card, error)
	ListByType(input string) ([]model.Card, error)
	GetByCpf(input string) (*model.Card, error)
	Create(input *model.Card) error
	Delete(id string) error
}

type CardUseCaseInterface interface {
	List() ([]model.CardResponse, error)
	ListByType(input string) ([]model.CardResponse, error)
	GetByCpf(input string) (*model.CardResponse, error)
	Create(input *model.CardRequest) (*model.CardResponse, error)
	Delete(id string) error
}
