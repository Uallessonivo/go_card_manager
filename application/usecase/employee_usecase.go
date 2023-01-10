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
	cards, _ := utils.CardResponse(cardsFound)

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

	var cards []*model.Card
	for _, employee := range employees {
		card, cardsErr := e.CardRepository.ListByOwner(employee.Cpf)
		if cardsErr != nil {
			return nil, cardsErr
		}
		cards = append(cards, card...)
	}

	results, er := utils.EmployeeResponse(employees, cards)
	if er != nil {
		return nil, er
	}

	return results, nil
}

func (e EmployeeUseCase) GetFiltered(input string) (*model.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeUseCase) UpdateEmployee(input *model.EmployeeRequest) (*model.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeUseCase) DeleteEmployee(input string) error {
	if er := e.EmployeeRepository.Delete(input); er != nil {
		return er
	}
	return nil
}
