package domain

import "context"

type Card struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Serial string `json:"serial"`
}

type CardRepository interface {
	List() ([]Card, error)
	ListByType(ctx context.Context, input string) ([]Card, error)
	GetByCpf(ctx context.Context, input string) (Card, error)
	GetByID(ctx context.Context, input string) (Card, error)
	Create(ctx context.Context, input *Card) error
	Update(ctx context.Context, input *Card) error
	Delete(ctx context.Context, id string) error
}
