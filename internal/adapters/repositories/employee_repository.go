package repositories

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"gorm.io/gorm"
)

type EmployeeRepository struct {
	Db *gorm.DB
}

func NewEmployeeRepository(Db *gorm.DB) ports.EmployeeRepository {
	return &EmployeeRepository{Db}
}

func (e EmployeeRepository) List() ([]*models.Employee, error) {
	var employees []*models.Employee
	if err := e.Db.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (e EmployeeRepository) Get(input string) (*models.Employee, error) {
	var employee *models.Employee
	if err := e.Db.Where("id = ? OR cpf = ?", input, input).First(&employee).Error; err != nil {
		return nil, err
	}
	return employee, nil
}

func (e EmployeeRepository) Create(input *models.Employee) error {
	if err := e.Db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Update(input *models.Employee) error {
	err := e.Db.Model(models.Employee{}).Where("id = ?", input.ID).UpdateColumns(input).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EmployeeRepository) Delete(id string) error {
	if err := e.Db.Delete(&models.Employee{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
