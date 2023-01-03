package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func CardsResponses(data []*model.Card) ([]*model.CardResponse, error) {
	var cards []*model.CardResponse

	if len(data) == 0 {
		return nil, errors.NoDataFound
	}

	for _, data := range data {
		response := model.CardResponse{
			ID:     data.ID,
			Type:   data.Type,
			Owner:  data.Owner,
			Serial: data.Serial,
		}
		cards = append(cards, &response)
	}
	return cards, nil
}
