package ports

import (
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type CardService interface {
	CreateCard(input *entities2.CardRequest) (*entities2.CardResponse, error)
	DeleteCard(id string) error
	ValidateCard(input string) (*entities2.Employee, error)
	ListAllCards() ([]*entities2.CardResponse, error)
	ListAllCardsByType(input string) ([]*entities2.CardResponse, error)
	ListAllCardsByOwner(input string) ([]*entities2.CardResponse, error)
}

type CardRepository interface {
	Create(input *entities2.Card) error
	Delete(id string) error
	List() ([]*entities2.Card, error)
	ListByTYpe(input string) ([]*entities2.Card, error)
	ListByOwner(input string) ([]*entities2.Card, error)
	GetByOwner(input string) (*entities2.Card, error)
	GetById(input string) (*entities2.Card, error)
}
