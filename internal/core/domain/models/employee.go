package models

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"regexp"

	"github.com/paemuri/brdoc"
	uuid "github.com/satori/go.uuid"
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
	matched := brdoc.IsCPF(cpf)
	return matched
}

func MakeEmployee(employee *EmployeeRequest) (*Employee, error) {
	if !validateName(employee.Name) {
		return nil, errors.InvalidFields
	}

	if !validateCpf(employee.Cpf) {
		return nil, errors.InvalidFields
	}

	return &Employee{
		ID:   uuid.NewV4().String(),
		Name: employee.Name,
		Cpf:  employee.Cpf,
	}, nil
}

func ValidateEmployee(employee *EmployeeRequest) error {
	if !validateName(employee.Name) {
		return errors.InvalidFields
	}

	if !validateCpf(employee.Cpf) {
		return errors.InvalidFields
	}

	return nil
}
