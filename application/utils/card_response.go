package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func CardResponse(data []*model.Card) []*model.CardResponse {
	var cards []*model.CardResponse
	for _, data := range data {
		cards = append(cards, &model.CardResponse{
			ID:     data.ID,
			Type:   data.Type,
			Owner:  data.Owner,
			Serial: data.Serial,
		})
	}
	return cards
}
