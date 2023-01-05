package usecase

import (
	"github.com/Uallessonivo/go_card_manager/application/utils"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type EmployeeUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
}

func NewEmployeeUseCase(u interfaces.EmployeeRepositoryInterface) interfaces.EmployeeUseCaseInterface {
	return &EmployeeUseCase{
		EmployeeRepository: u,
	}
}

func (e EmployeeUseCase) CreateEmployee(input *model.EmployeeRequest) (*model.EmployeeResponse, error) {
	newEmployee, err := model.MakeEmployee(input)
	if err != nil {
		return nil, err
	}

	if er := e.EmployeeRepository.Create(newEmployee); er != nil {
		return nil, er
	}

	cards, errr := utils.CardResponse(newEmployee.Cards)
	if errr != nil {
		return nil, errr
	}

	return &model.EmployeeResponse{
		ID:    newEmployee.ID,
		Name:  newEmployee.Name,
		Cpf:   newEmployee.Cpf,
		Cards: cards,
	}, nil
}

func (e EmployeeUseCase) ListEmployees() ([]*model.EmployeeResponse, error) {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}
