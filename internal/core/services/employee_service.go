package services

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
)

type EmployeeUseCase struct {
	EmployeeRepository ports.EmployeeRepository
	CardRepository     ports.CardRepository
}

func NewEmployeeService(u ports.EmployeeRepository, c ports.CardRepository) ports.EmployeeService {
	return &EmployeeUseCase{
		EmployeeRepository: u,
		CardRepository:     c,
	}
}

func (e EmployeeUseCase) CreateEmployee(input *models.EmployeeRequest) (*models.EmployeeResponse, error) {
	newEmployee, err := models.MakeEmployee(input)
	if err != nil {
		return nil, err
	}

	employeeExists, _ := e.EmployeeRepository.Get(newEmployee.Cpf)
	if employeeExists != nil {
		return nil, errors.AlreadyExists
	}

	if er := e.EmployeeRepository.Create(newEmployee); er != nil {
		return nil, er
	}

	cardsFound, _ := e.CardRepository.ListByOwner(newEmployee.Cpf)
	cards := utils.CardResponse(cardsFound)

	return &models.EmployeeResponse{
		ID:    newEmployee.ID,
		Name:  newEmployee.Name,
		Cpf:   newEmployee.Cpf,
		Cards: cards,
	}, nil
}

func (e EmployeeUseCase) ListEmployees() ([]*models.EmployeeResponse, error) {
	employees, err := e.EmployeeRepository.List()
	if err != nil {
		return nil, err
	}

	cards, _ := e.CardRepository.List()

	employeeResponses := utils.EmployeeResponse(employees, cards)

	return employeeResponses, nil
}

func (e EmployeeUseCase) GetFiltered(input string) (*models.EmployeeResponse, error) {
	employeeFound, err := e.EmployeeRepository.Get(input)

	if err != nil {
		return nil, errors.NotFound
	}

	cardsFound, _ := e.CardRepository.ListByOwner(employeeFound.Cpf)
	cards := utils.CardResponse(cardsFound)

	return &models.EmployeeResponse{
		ID:    employeeFound.ID,
		Name:  employeeFound.Name,
		Cpf:   employeeFound.Cpf,
		Cards: cards,
	}, nil
}

func (e EmployeeUseCase) UpdateEmployee(id string, input *models.EmployeeRequest) (*models.EmployeeResponse, error) {
	_, err := e.EmployeeRepository.Get(id)
	if err != nil {
		return nil, errors.NotFound
	}

	if validateErr := models.ValidateEmployee(input); validateErr != nil {
		return nil, errors.InvalidFields
	}

	if updateErr := e.EmployeeRepository.Update(&models.Employee{
		ID:   id,
		Name: input.Name,
		Cpf:  input.Cpf,
	}); updateErr != nil {
		return nil, updateErr
	}

	return &models.EmployeeResponse{
		ID:   id,
		Name: input.Name,
		Cpf:  input.Cpf,
	}, nil
}

func (e EmployeeUseCase) DeleteEmployee(input string) error {
	if er := e.EmployeeRepository.Delete(input); er != nil {
		return er
	}
	return nil
}
