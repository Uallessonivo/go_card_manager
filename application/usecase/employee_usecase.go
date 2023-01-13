package usecase

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type EmployeeUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardRepository     interfaces.CardRepositoryInterface
}

func NewEmployeeUseCase(u interfaces.EmployeeRepositoryInterface, c interfaces.CardRepositoryInterface) interfaces.EmployeeUseCaseInterface {
	return &EmployeeUseCase{
		EmployeeRepository: u,
		CardRepository:     c,
	}
}

func (e EmployeeUseCase) CreateEmployee(input *model.EmployeeRequest) (*model.EmployeeResponse, error) {
	newEmployee, err := model.MakeEmployee(input)
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

	return &model.EmployeeResponse{
		ID:    newEmployee.ID,
		Name:  newEmployee.Name,
		Cpf:   newEmployee.Cpf,
		Cards: cards,
	}, nil
}

func (e EmployeeUseCase) ListEmployees() ([]*model.EmployeeResponse, error) {
	employees, err := e.EmployeeRepository.List()
	if err != nil {
		return nil, err
	}

	cards, _ := e.CardRepository.List()

	employeeResponses := utils.EmployeeResponse(employees, cards)

	return employeeResponses, nil
}

func (e EmployeeUseCase) GetFiltered(input string) (*model.EmployeeResponse, error) {
	employeeFound, err := e.EmployeeRepository.Get(input)

	if err != nil {
		return nil, errors.NotFound
	}

	cardsFound, _ := e.CardRepository.ListByOwner(employeeFound.Cpf)
	cards := utils.CardResponse(cardsFound)

	return &model.EmployeeResponse{
		ID:    employeeFound.ID,
		Name:  employeeFound.Name,
		Cpf:   employeeFound.Cpf,
		Cards: cards,
	}, nil
}

func (e EmployeeUseCase) UpdateEmployee(id string, input *model.EmployeeRequest) (*model.EmployeeResponse, error) {
	_, err := e.EmployeeRepository.Get(id)
	if err != nil {
		return nil, errors.NotFound
	}

	if validateErr := model.ValidateEmployee(input); validateErr != nil {
		return nil, errors.InvalidFields
	}

	if updateErr := e.EmployeeRepository.Update(&model.Employee{
		ID:   id,
		Name: input.Name,
		Cpf:  input.Cpf,
	}); updateErr != nil {
		return nil, updateErr
	}

	return &model.EmployeeResponse{
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
