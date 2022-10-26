package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type EmployeeRepositoryInterface interface {
	List() ([]*model.Employee, error)
	GetByCpf(input string) (*model.Employee, error)
	GetByName(input string) (*model.Employee, error)
	Create(input *model.Employee) error
	Update(input *model.Employee) error
	Delete(id string) error
}
