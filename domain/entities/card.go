package entities

import (
	"github.com/Uallessonivo/go_card_manager/domain/enums"
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
	Type   enums.CardType `json:"type"`
	Owner  string         `json:"owner"`
	Serial string         `json:"serial"`
}

type CardResponse struct {
	ID     string `json:"id"`
	Type   string `json:"type"`
	Owner  string `json:"owner"`
	Serial string `json:"serial"`
}

func MakeCard(card *CardRequest, ownerName string) (*Card, error) {
	validTypes := map[enums.CardType]bool{
		enums.DespesasMatriz: true,
		enums.DespesasFilial: true,
		enums.IncentivoCap1:  true,
		enums.IncentivoCap2:  true,
		enums.IncentivoInt1:  true,
		enums.IncentivoInt2:  true,
		enums.IncentivoSe:    true,
	}

	if !validTypes[card.Type] {
		return nil, errors.InvalidFields
	}

	if len(card.Serial) != 15 || len(card.Owner) != 11 {
		return nil, errors.InvalidFields
	}

	return &Card{
		ID:     uuid.NewV4().String(),
		Type:   string(card.Type),
		Owner:  card.Owner,
		Name:   ownerName,
		Serial: card.Serial,
	}, nil
}
