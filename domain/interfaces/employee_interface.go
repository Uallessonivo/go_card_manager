package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/entities"

type EmployeeUseCaseInterface interface {
	CreateEmployee(input *entities.EmployeeRequest) (*entities.EmployeeResponse, error)
	ListEmployees() ([]*entities.EmployeeResponse, error)
	GetFiltered(input string) (*entities.EmployeeResponse, error)
	UpdateEmployee(id string, input *entities.EmployeeRequest) (*entities.EmployeeResponse, error)
	DeleteEmployee(input string) error
}

type EmployeeRepositoryInterface interface {
	List() ([]*entities.Employee, error)
	Get(input string) (*entities.Employee, error)
	Create(input *entities.Employee) error
	Update(input *entities.Employee) error
	Delete(id string) error
}
