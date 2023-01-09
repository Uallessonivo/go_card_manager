package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func EmployeeResponse(data []*model.Employee) ([]*model.EmployeeResponse, error) {
	if len(data) == 0 {
		return nil, errors.NoDataFound
	}

	var employees []*model.EmployeeResponse
	for _, data := range data {
		card, err := CardResponse(data.Cards)
		if err != nil {
			return nil, err
		}
		employees = append(employees, &model.EmployeeResponse{
			ID:    data.ID,
			Name:  data.Name,
			Cpf:   data.Cpf,
			Cards: card,
		})
	}
	return employees, nil
}
