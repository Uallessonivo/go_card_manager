package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	Db *gorm.DB
}

func NewEmployeeRepository(Db *gorm.DB) interfaces.EmployeeRepositoryInterface {
	return &EmployeeRepository{Db}
}

func (e EmployeeRepository) List() ([]*model.Employee, error) {
	var employees []*model.Employee
	if err := e.Db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (e EmployeeRepository) Get(input string) (*model.Employee, error) {
	var employee *model.Employee
	if err := e.Db.Where("id = ? OR cpf = ?", input, input).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (e EmployeeRepository) Create(input *model.Employee) error {
	if err := e.Db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Update(input *model.Employee) error {
	err := e.Db.Model(model.Employee{}).Where("id = ?", input.ID).UpdateColumns(input).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Delete(id string) error {
	if err := e.Db.Delete(&model.Employee{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
