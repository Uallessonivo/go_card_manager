package model

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	uuid "github.com/satori/go.uuid"
	"regexp"
)

type Employee struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Cpf   string
	Cards []*Card `gorm:"many2many:cards;"`
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
	if len(employee.Cpf) != 11 {
		return nil, errors.InvalidFields
	}

	if !validateName(employee.Name) {
		return nil, errors.InvalidFields
	}

	var cards []*Card
	if employee.Cards != nil {
		for _, data := range employee.Cards {
			cards = append(cards, &Card{
				ID:     uuid.NewV4().String(),
				Type:   string(data.Type),
				Owner:  data.Owner,
				Name:   data.Name,
				Serial: data.Serial,
			})
		}
	}

	return &Employee{
		ID:    uuid.NewV4().String(),
		Name:  employee.Name,
		Cpf:   employee.Cpf,
		Cards: cards,
	}, nil
}
