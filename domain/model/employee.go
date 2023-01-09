package model

import (
	uuid "github.com/satori/go.uuid"
	"regexp"
)

type Employee struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Cpf   string
	Cards []*Card `gorm:"many2many:cards"`
}

type EmployeeRequest struct {
	Name  string         `json:"name"`
	Cpf   string         `json:"cpf"`
	Cards []*CardRequest `json:"cards"`
}

type EmployeeResponse struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Cpf   string          `json:"cpf"`
	Cards []*CardResponse `json:"cards"`
}

func validateName(name string) bool {
	matched, _ := regexp.MatchString(`\b([a-zÀ-ÿA-Z][-,a-z. ']+[ ]*)+`, name)
	return matched
}

func MakeEmployee(employee *EmployeeRequest) (*Employee, error) {
	var cards []*Card
	if employee.Cards != nil {
		for _, data := range employee.Cards {
			card, err := MakeCard(data)
			if err != nil {
				return nil, err
			}
			cards = append(cards, card)
		}
	}

	return &Employee{
		ID:    uuid.NewV4().String(),
		Name:  employee.Name,
		Cpf:   employee.Cpf,
		Cards: cards,
	}, nil
}
