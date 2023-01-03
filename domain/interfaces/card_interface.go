package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type CardUseCaseInterface interface {
	ListAllCards() ([]*model.CardResponse, error)
	ListAllCardsByType(input string) ([]*model.CardResponse, error)
	ListAllCardsByOwner(input string) ([]*model.CardResponse, error)
	CreateCard(input *model.CardRequest) (*model.CardResponse, error)
	DeleteCard(id string) error
}

type CardRepositoryInterface interface {
	List() ([]*model.Card, error)
	ListByTYpe(input string) ([]*model.Card, error)
	ListByOwner(input string) ([]*model.Card, error)
	GetById(input string) (*model.Card, error)
	Create(input *model.Card) error
	Delete(id string) error
}
