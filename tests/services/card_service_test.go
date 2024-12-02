package services_test

import (
	"testing"

	"github.com/Uallessonivo/go_card_manager/internal/core/domain/enums"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCardRepository struct {
	mock.Mock
}

type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) List() ([]*models.Employee, error) {
	args := m.Called()
	return args.Get(0).([]*models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Update(input *models.Employee) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockCardRepository) List() ([]*models.Card, error) {
	args := m.Called()
	return args.Get(0).([]*models.Card), args.Error(1)
}

func (m *MockCardRepository) ListByTYpe(input string) ([]*models.Card, error) {
	args := m.Called(input)
	return args.Get(0).([]*models.Card), args.Error(1)
}

func (m *MockCardRepository) ListByOwner(owner string) ([]*models.Card, error) {
	args := m.Called(owner)
	return args.Get(0).([]*models.Card), args.Error(1)
}

func (m *MockCardRepository) GetByOwner(owner string) (*models.Card, error) {
	args := m.Called(owner)
	return args.Get(0).(*models.Card), args.Error(1)
}

func (m *MockCardRepository) Create(card *models.Card) error {
	args := m.Called(card)
	return args.Error(0)
}

func (m *MockCardRepository) GetById(id string) (*models.Card, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Card), args.Error(1)
}

func (m *MockCardRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockEmployeeRepository) Get(id string) (*models.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) Create(employee *models.Employee) error {
	args := m.Called(employee)
	return args.Error(0)
}

func (m *MockEmployeeRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCard_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	input := &models.CardRequest{
		Type:   enums.CardType("DESPESAS_MATRIZ"),
		Owner:  "00000000000",
		Serial: "000000000000000",
	}

	employee := &models.Employee{
		ID:   "1",
		Name: "John",
		Cpf:  "00000000000",
	}

	// Mock para quando o Employee for buscado
	mockEmployeeRepo.On("Get", "00000000000").Return(employee, nil)
	// Mock para quando a lista de cartões do proprietário for buscada
	mockCardRepo.On("ListByOwner", "00000000000").Return([]*models.Card{}, nil)
	// Mock para quando o método Create for chamado
	mockCardRepo.On("Create", mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		// Verifique se o novo card criado tem os valores corretos
		card := args.Get(0).(*models.Card)
		assert.Equal(t, "DESPESAS_MATRIZ", card.Type)
		assert.Equal(t, input.Owner, card.Owner)
		assert.Equal(t, input.Serial, card.Serial)
	})

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.CreateCard(input)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "DESPESAS_MATRIZ", result.Type)
	assert.Equal(t, input.Owner, result.Owner)
	assert.Equal(t, input.Serial, result.Serial)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}
