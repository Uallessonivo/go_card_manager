package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	Db *gorm.DB
}

func NewEmployeeRepository(Db *gorm.DB) interfaces.EmployeeRepositoryInterface {
	return &EmployeeRepository{Db}
}

func (e EmployeeRepository) List() ([]*entities.Employee, error) {
	var employees []*entities.Employee
	if err := e.Db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (e EmployeeRepository) Get(input string) (*entities.Employee, error) {
	var employee *entities.Employee
	if err := e.Db.Where("id = ? OR cpf = ?", input, input).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (e EmployeeRepository) Create(input *entities.Employee) error {
	if err := e.Db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Update(input *entities.Employee) error {
	err := e.Db.Model(entities.Employee{}).Where("id = ?", input.ID).UpdateColumns(input).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Delete(id string) error {
	if err := e.Db.Delete(&entities.Employee{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
