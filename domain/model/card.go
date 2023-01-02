package model

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	uuid "github.com/satori/go.uuid"
)

type Card struct {
	ID     string `gorm:"primary_key"`
	Type   string
	Owner  string
	Name   string
	Serial string
}

type CardRequest struct {
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Serial string `json:"serial"`
}

type CardResponse struct {
	ID     string
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Serial string `json:"serial"`
}

func MakeCard(card *CardRequest) (*Card, error) {
	if len(card.Serial) != 15 || len(card.Owner) != 11 {
		return nil, errors.InvalidFields
	}

	newCard := Card{
		ID:     uuid.NewV4().String(),
		Type:   card.Type,
		Owner:  card.Owner,
		Name:   card.Name,
		Serial: card.Serial,
	}

	return &newCard, nil
}
