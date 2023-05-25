package ports

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type CardService interface {
	CreateCard(input *models.CardRequest) (*models.CardResponse, error)
	DeleteCard(id string) error
	ValidateCard(input string) (*models.Employee, error)
	ListAllCards() ([]*models.CardResponse, error)
	ListAllCardsByType(input string) ([]*models.CardResponse, error)
	ListAllCardsByOwner(input string) ([]*models.CardResponse, error)
}

type CardRepository interface {
	Create(input *models.Card) error
	Delete(id string) error
	List() ([]*models.Card, error)
	ListByTYpe(input string) ([]*models.Card, error)
	ListByOwner(input string) ([]*models.Card, error)
	GetByOwner(input string) (*models.Card, error)
	GetById(input string) (*models.Card, error)
}
