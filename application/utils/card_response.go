package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
)

func CardResponse(data []*entities.Card) []*entities.CardResponse {
	var cards []*entities.CardResponse
	for _, data := range data {
		cards = append(cards, &entities.CardResponse{
			ID:     data.ID,
			Type:   data.Type,
			Owner:  data.Owner,
			Serial: data.Serial,
		})
	}
	return cards
}
