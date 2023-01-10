package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func EmployeeResponse(dataEmployees []*model.Employee, dataCards []*model.Card) ([]*model.EmployeeResponse, error) {
	var cards []*model.CardResponse
	if dataCards != nil {
		for _, data := range dataCards {
			card := &model.CardResponse{
				ID:     data.ID,
				Type:   data.Type,
				Owner:  data.Owner,
				Serial: data.Serial,
			}
			cards = append(cards, card)
		}
	}

	var employees []*model.EmployeeResponse
	for _, data := range dataEmployees {
		employees = append(employees, &model.EmployeeResponse{
			ID:    data.ID,
			Name:  data.Name,
			Cpf:   data.Cpf,
			Cards: cards,
		})
	}

	return employees, nil
}
