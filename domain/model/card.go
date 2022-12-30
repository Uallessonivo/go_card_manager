package model

import uuid "github.com/satori/go.uuid"

type Card struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Name   string `json:"name"`
	Serial string `json:"serial"`
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
	// TODO: validate fields
	newCard := Card{
		ID:     uuid.NewV4().String(),
		Type:   card.Type,
		Owner:  card.Owner,
		Name:   card.Name,
		Serial: card.Serial,
	}

	return &newCard, nil
}
