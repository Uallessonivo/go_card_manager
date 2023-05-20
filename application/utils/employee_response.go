package utils

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
)

func EmployeeResponse(dataEmployees []*entities.Employee, dataCards []*entities.Card) []*entities.EmployeeResponse {
	employeeCards := make(map[string][]*entities.Card)

	for _, card := range dataCards {
		employeeCards[card.Owner] = append(employeeCards[card.Owner], card)
	}

	var employeeResponses []*entities.EmployeeResponse

	for _, employee := range dataEmployees {
		employeeResponse := &entities.EmployeeResponse{
			ID:   employee.ID,
			Name: employee.Name,
			Cpf:  employee.Cpf,
		}

		if dataCards, ok := employeeCards[employee.Cpf]; ok {
			employeeResponse.Cards = CardResponse(dataCards)
		}

		employeeResponses = append(employeeResponses, employeeResponse)
	}

	return employeeResponses
}
