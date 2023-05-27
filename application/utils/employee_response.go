package utils

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

func EmployeeResponse(dataEmployees []*models.Employee, dataCards []*models.Card) []*models.EmployeeResponse {
	employeeCards := make(map[string][]*models.Card)

	for _, card := range dataCards {
		employeeCards[card.Owner] = append(employeeCards[card.Owner], card)
	}

	var employeeResponses []*models.EmployeeResponse

	for _, employee := range dataEmployees {
		employeeResponse := &models.EmployeeResponse{
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
