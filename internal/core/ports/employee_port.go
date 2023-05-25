package ports

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type EmployeeService interface {
	CreateEmployee(input *models.EmployeeRequest) (*models.EmployeeResponse, error)
	ListEmployees() ([]*models.EmployeeResponse, error)
	GetFiltered(input string) (*models.EmployeeResponse, error)
	UpdateEmployee(id string, input *models.EmployeeRequest) (*models.EmployeeResponse, error)
	DeleteEmployee(input string) error
}

type EmployeeRepository interface {
	List() ([]*models.Employee, error)
	Get(input string) (*models.Employee, error)
	Create(input *models.Employee) error
	Update(input *models.Employee) error
	Delete(id string) error
}
