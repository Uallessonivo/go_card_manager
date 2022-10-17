package model

type Card struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Serial string `json:"serial"`
}

type CardRepository interface {
	List() ([]*Card, error)
	ListByType(input string) ([]*Card, error)
	GetByCpf(input string) (*Card, error)
	GetByID(input string) (*Card, error)
	Create(input *Card) error
	Update(input *Card) error
	Delete(id string) error
}
