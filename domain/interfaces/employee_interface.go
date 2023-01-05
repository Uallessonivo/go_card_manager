package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type EmployeeUseCaseInterface interface {
	CreateEmployee(input *model.EmployeeRequest) (*model.EmployeeResponse, error)
	ListEmployees() ([]*model.EmployeeResponse, error)
	GetFiltered(input string) (*model.EmployeeResponse, error)
	UpdateEmployee(input *model.EmployeeRequest) (*model.EmployeeResponse, error)
	DeleteEmployee(input string) error
}

type EmployeeRepositoryInterface interface {
	List() ([]*model.Employee, error)
	Get(input string) (*model.Employee, error)
	Create(input *model.Employee) error
	Update(input *model.Employee) error
	Delete(id string) error
}
