package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

func EmployeeResponse(dataEmployees []*model.Employee, dataCards []*model.Card) ([]*model.EmployeeResponse, error) {
	employeeCards := make(map[string][]*model.Card)

	for _, card := range dataCards {
		employeeCards[card.Owner] = append(employeeCards[card.Owner], card)
	}

	var employeeResponses []*model.EmployeeResponse

	for _, employee := range dataEmployees {
		employeeResponse := &model.EmployeeResponse{
			ID:   employee.ID,
			Name: employee.Name,
			Cpf:  employee.Cpf,
		}

		if dataCards, ok := employeeCards[employee.Cpf]; ok {
			employeeResponse.Cards = CardResponse(dataCards)
		}

		employeeResponses = append(employeeResponses, employeeResponse)
	}

	return employeeResponses, nil
}
