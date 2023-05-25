package utils

import (
	entities2 "github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

func EmployeeResponse(dataEmployees []*entities2.Employee, dataCards []*entities2.Card) []*entities2.EmployeeResponse {
	employeeCards := make(map[string][]*entities2.Card)

	for _, card := range dataCards {
		employeeCards[card.Owner] = append(employeeCards[card.Owner], card)
	}

	var employeeResponses []*entities2.EmployeeResponse

	for _, employee := range dataEmployees {
		employeeResponse := &entities2.EmployeeResponse{
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
