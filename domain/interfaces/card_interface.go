package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/entities"

type CardUseCaseInterface interface {
	CreateCard(input *entities.CardRequest) (*entities.CardResponse, error)
	DeleteCard(id string) error
	ValidateCard(input string) (*entities.Employee, error)
	ListAllCards() ([]*entities.CardResponse, error)
	ListAllCardsByType(input string) ([]*entities.CardResponse, error)
	ListAllCardsByOwner(input string) ([]*entities.CardResponse, error)
}

type CardRepositoryInterface interface {
	Create(input *entities.Card) error
	Delete(id string) error
	List() ([]*entities.Card, error)
	ListByTYpe(input string) ([]*entities.Card, error)
	ListByOwner(input string) ([]*entities.Card, error)
	GetByOwner(input string) (*entities.Card, error)
	GetById(input string) (*entities.Card, error)
}
