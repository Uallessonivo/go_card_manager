package model

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	uuid "github.com/satori/go.uuid"
	"regexp"
)

type Employee struct {
	ID   string `gorm:"primary_key"`
	Name string
	Cpf  string
}

type EmployeeRequest struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
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

func validateCpf(cpf string) bool {
	matched, _ := regexp.MatchString("", cpf)
	return matched
}

func MakeEmployee(employee *EmployeeRequest) (*Employee, error) {
	if !validateName(employee.Name) {
		return nil, errors.InvalidFields
	}

	// TODO
	//if !validateCpf(employee.Cpf) {
	//	return nil, errors.InvalidFields
	//}

	return &Employee{
		ID:   uuid.NewV4().String(),
		Name: employee.Name,
		Cpf:  employee.Cpf,
	}, nil
}
