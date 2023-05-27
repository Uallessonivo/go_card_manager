package utils

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

func CardResponse(data []*models.Card) []*models.CardResponse {
	var cards []*models.CardResponse
	for _, data := range data {
		cards = append(cards, &models.CardResponse{
			ID:     data.ID,
			Type:   data.Type,
			Owner:  data.Owner,
			Serial: data.Serial,
		})
	}
	return cards
}
