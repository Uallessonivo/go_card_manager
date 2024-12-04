package services_test

import (
	"testing"

	errs "github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateEmployee_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.EmployeeRequest{
		Name: "John",
		Cpf:  "72345279079", // Fake valid CPF
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employee.Cpf).Return((*models.Employee)(nil), nil)
	// Mock para quando o employee for criado
	mockEmployeeRepo.On("Create", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		// Verifique se o novo employee criado tem os valores corretos
		empl := args.Get(0).(*models.Employee)
		assert.Equal(t, employee.Cpf, empl.Cpf)
		assert.Equal(t, employee.Name, empl.Name)
	})

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.CreateEmployee(employee)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestCreateEmployee_AlreadyExists(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.EmployeeRequest{
		Name: "John",
		Cpf:  "72345279079", // Fake valid CPF
	}

	existingEmployee := &models.Employee{
		ID:   "1",
		Name: "John",
		Cpf:  "72345279079",
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employee.Cpf).Return(existingEmployee, nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.CreateEmployee(employee)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.AlreadyExists, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestCreateEmployee_InvalidCpfError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.EmployeeRequest{
		Name: "John",
		Cpf:  "00000000000",
	}

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.CreateEmployee(employee)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, errs.InvalidCpf, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestCreateEmployee_InvalidNameError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.EmployeeRequest{
		Name: "John 21",
		Cpf:  "72345279079", // Fake valid cpf
	}

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.CreateEmployee(employee)

	// Assert
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, errs.InvalidName, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}
