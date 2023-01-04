package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func CardsResponses(data []*model.Card) ([]*model.CardResponse, error) {
	if len(data) == 0 {
		return nil, errors.NoDataFound
	}

	var cards []*model.CardResponse
	for _, data := range data {
		cards = append(cards, &model.CardResponse{
			ID:     data.ID,
			Type:   data.Type,
			Owner:  data.Owner,
			Serial: data.Serial,
		})
	}
	return cards, nil
}
