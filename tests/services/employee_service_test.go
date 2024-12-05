package services_test

import (
	"errors"
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

func TestListEmployees_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employees := []*models.Employee{
		{
			ID:   "1",
			Name: "John",
			Cpf:  "72345279079",
		},
		{
			ID:   "2",
			Name: "Jane",
			Cpf:  "12345678901",
		},
	}

	cards := []*models.Card{
		{
			ID:     "1",
			Owner:  "72345279079",
			Serial: "000000000000000",
		},
		{
			ID:     "2",
			Owner:  "12345678901",
			Serial: "000000000000000",
		},
	}

	// Mock para quando os employees forem listados
	mockEmployeeRepo.On("List").Return(employees, nil)
	// Mock para quando os cards forem listados
	mockCardRepo.On("List").Return(cards, nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.ListEmployees()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestListEmployees_Error(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	// Mock para quando os employees forem listados e ocorrer um erro
	mockEmployeeRepo.On("List").Return([]*models.Employee(nil), errors.New("some error"))

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.ListEmployees()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestGetFiltered_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.Employee{
		ID:   "1",
		Name: "John",
		Cpf:  "72345279079",
	}

	cards := []*models.Card{
		{
			ID:     "1",
			Owner:  "72345279079",
			Serial: "000000000000000",
		},
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employee.Cpf).Return(employee, nil)
	// Mock para quando os cards forem listados pelo dono
	mockCardRepo.On("ListByOwner", employee.Cpf).Return(cards, nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.GetFiltered(employee.Cpf)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, employee.ID, result.ID)
	assert.Equal(t, employee.Name, result.Name)
	assert.Equal(t, employee.Cpf, result.Cpf)
	assert.Len(t, result.Cards, 1)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestGetFiltered_NotFoundError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	// Mock para quando o employee não for encontrado
	mockEmployeeRepo.On("Get", "72345279079").Return((*models.Employee)(nil), errors.New("not found"))

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.GetFiltered("72345279079")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.NotFound, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestGetFiltered_NoCardsSuccess(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employee := &models.Employee{
		ID:   "1",
		Name: "John",
		Cpf:  "72345279079",
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employee.Cpf).Return(employee, nil)
	// Mock para quando não houver cards para o dono
	mockCardRepo.On("ListByOwner", employee.Cpf).Return([]*models.Card{}, nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.GetFiltered(employee.Cpf)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, employee.ID, result.ID)
	assert.Equal(t, employee.Name, result.Name)
	assert.Equal(t, employee.Cpf, result.Cpf)
	assert.Len(t, result.Cards, 0)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestUpdateEmployee_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"
	employeeRequest := &models.EmployeeRequest{
		Name: "John Updated",
		Cpf:  "72345279079", // Fake valid CPF
	}

	existingEmployee := &models.Employee{
		ID:   employeeID,
		Name: "John",
		Cpf:  "72345279079",
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employeeID).Return(existingEmployee, nil)
	// Mock para quando o employee for atualizado
	mockEmployeeRepo.On("Update", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		// Verifique se o employee atualizado tem os valores corretos
		empl := args.Get(0).(*models.Employee)
		assert.Equal(t, employeeRequest.Cpf, empl.Cpf)
		assert.Equal(t, employeeRequest.Name, empl.Name)
		assert.Equal(t, employeeID, empl.ID)
	})

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.UpdateEmployee(employeeID, employeeRequest)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, employeeRequest.Name, result.Name)
	assert.Equal(t, employeeRequest.Cpf, result.Cpf)
	assert.Equal(t, employeeID, result.ID)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestUpdateEmployee_NotFound(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"
	employeeRequest := &models.EmployeeRequest{
		Name: "John Updated",
		Cpf:  "72345279079", // Fake valid CPF
	}

	// Mock para quando o employee não for encontrado
	mockEmployeeRepo.On("Get", employeeID).Return((*models.Employee)(nil), errors.New("some error"))

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.UpdateEmployee(employeeID, employeeRequest)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.NotFound, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestUpdateEmployee_InvalidFields(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"
	employeeRequest := &models.EmployeeRequest{
		Name: "John 21",
		Cpf:  "00000000000", // Invalid CPF
	}

	existingEmployee := &models.Employee{
		ID:   employeeID,
		Name: "John",
		Cpf:  "72345279079",
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employeeID).Return(existingEmployee, nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.UpdateEmployee(employeeID, employeeRequest)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.InvalidFields, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestUpdateEmployee_UpdateError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"
	employeeRequest := &models.EmployeeRequest{
		Name: "John Updated",
		Cpf:  "72345279079", // Fake valid CPF
	}

	existingEmployee := &models.Employee{
		ID:   employeeID,
		Name: "John",
		Cpf:  "72345279079",
	}

	// Mock para quando o employee for buscado
	mockEmployeeRepo.On("Get", employeeID).Return(existingEmployee, nil)
	// Mock para quando ocorrer um erro ao atualizar o employee
	mockEmployeeRepo.On("Update", mock.Anything).Return(errors.New("update error"))

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	result, err := service.UpdateEmployee(employeeID, employeeRequest)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "update error", err.Error())

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestDeleteEmployee_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"

	// Mock para quando o employee for deletado
	mockEmployeeRepo.On("Delete", employeeID).Return(nil)

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	err := service.DeleteEmployee(employeeID)

	// Assert
	assert.NoError(t, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestDeleteEmployee_Error(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	employeeID := "1"

	// Mock para quando ocorrer um erro ao deletar o employee
	mockEmployeeRepo.On("Delete", employeeID).Return(errors.New("delete error"))

	service := services.NewEmployeeService(mockEmployeeRepo, mockCardRepo)

	// Act
	err := service.DeleteEmployee(employeeID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "delete error", err.Error())

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}


