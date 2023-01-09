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
	if err := e.Db.Preload("Cards").Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (e EmployeeRepository) Get(input string) (*model.Employee, error) {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeRepository) Create(input *model.Employee) error {
	var employee model.Employee

	err := e.Db.Model(&employee).Association("Cards").Append(&input)
	if err != nil {
		return err
	}

	return nil
}

func (e EmployeeRepository) Update(input *model.Employee) error {
	//TODO implement me
	panic("implement me")
}

func (e EmployeeRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
