package services_test

import (
	"errors"
	"testing"

	"github.com/Uallessonivo/go_card_manager/internal/core/domain/enums"
	errs "github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
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

func (m *MockCardRepository) ListByType(input string) ([]*models.Card, error) {
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

func TestCreateCard_MaxNumberOfCardsError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	input := &models.CardRequest{
		Type:   enums.CardType("DESPESAS_MATRIZ"),
		Owner:  "00000000000",
		Serial: "222222222222222",
	}

	employee := &models.Employee{
		ID:   "1",
		Name: "John",
		Cpf:  "00000000000",
	}

	// Mock para quando o Employee for buscado
	mockEmployeeRepo.On("Get", "00000000000").Return(employee, nil)
	// Mock para quando a lista de cartões do proprietário for buscada
	mockCardRepo.On("ListByOwner", "00000000000").Return([]*models.Card{
		{ID: "1", Type: "DESPESAS_MATRIZ", Owner: "00000000000", Serial: "111111111111111"},
		{ID: "2", Type: "INCENTIVO_MATRIZ", Owner: "00000000000", Serial: "222222222222222"},
	}, nil)

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.CreateCard(input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.MaxNumberOfCards, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}

func TestCreateCard_OwnerNotFoundError(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	input := &models.CardRequest{
		Type:   enums.CardType("DESPESAS_MATRIZ"),
		Owner:  "00000000000",
		Serial: "222222222222222",
	}

	// Mock para quando o Employee for buscado
	mockEmployeeRepo.On("Get", "00000000000").Return((*models.Employee)(nil), errs.OwnerNotFound)

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.CreateCard(input)

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, errs.OwnerNotFound, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockEmployeeRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
	mockCardRepo.AssertNotCalled(t, "ListByOwner")
	mockCardRepo.AssertNotCalled(t, "Create")
}

func TestListAllCards_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	cards := []*models.Card{
		{
			ID:     "1",
			Type:   "DESPESAS_MATRIZ",
			Owner:  "00000000000",
			Serial: "000000000000000",
		},
		{
			ID:     "2",
			Type:   "DESPESAS_FILIAL",
			Owner:  "11111111111",
			Serial: "111111111111111",
		},
	}

	// Mock para quando a lista de cartões for buscada
	mockCardRepo.On("List").Return(cards, nil)

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.ListAllCards()

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "DESPESAS_MATRIZ", result[0].Type)
	assert.Equal(t, "00000000000", result[0].Owner)
	assert.Equal(t, "000000000000000", result[0].Serial)
	assert.Equal(t, "DESPESAS_FILIAL", result[1].Type)
	assert.Equal(t, "11111111111", result[1].Owner)
	assert.Equal(t, "111111111111111", result[1].Serial)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
}

func TestListAllCards_Error(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	// Mock para quando a lista de cartões for buscada e ocorrer um erro
	mockCardRepo.On("List").Return([]*models.Card(nil), errors.New("some error"))

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.ListAllCards()

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
}

func TestListAllCardsByType_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	cards := []*models.Card{
		{
			ID:     "1",
			Type:   "DESPESAS_MATRIZ",
			Owner:  "00000000000",
			Serial: "000000000000000",
		},
		{
			ID:     "2",
			Type:   "DESPESAS_MATRIZ",
			Owner:  "11111111111",
			Serial: "111111111111111",
		},
	}

	// Mock para quando a lista de cartões for buscada
	mockCardRepo.On("ListByType", "DESPESAS_MATRIZ").Return(cards, nil)

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.ListAllCardsByType("DESPESAS_MATRIZ")

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result, 2)
	assert.Equal(t, "DESPESAS_MATRIZ", result[0].Type)
	assert.Equal(t, "DESPESAS_MATRIZ", result[1].Type)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
}

func TestListAllCardsByType_Error(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	// Mock para quando a lista de cartões for buscada
	mockCardRepo.On("ListByType", "DESPESAS_MATRIZ").Return([]*models.Card(nil), errors.New("some error"))

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	result, err := service.ListAllCardsByType("DESPESAS_MATRIZ")

	// Assert
	assert.Error(t, err)
	assert.Nil(t, result)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
}

func TestDeleteCard_Success(t *testing.T) {
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	card := &models.Card{
		ID:     "1",
		Type:   "DESPESAS_MATRIZ",
		Owner:  "00000000000",
		Serial: "111111111111111",
	}

	// Mock para quando a lista de cartões for buscada
	mockCardRepo.On("GetById", "00000000000").Return(card, nil)
	mockCardRepo.On("Delete", "00000000000").Return(nil)

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	err := service.DeleteCard("00000000000")

	// Assert
	assert.NoError(t, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
}

func TestDeleteCard_UserNotFoundError(t *testing.T) {
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)

	// Mock para quando a lista de cartões for buscada
	mockCardRepo.On("GetById", "00000000000").Return((*models.Card)(nil), errors.New("some error"))

	service := services.NewCardService(mockCardRepo, mockEmployeeRepo)

	// Act
	err := service.DeleteCard("00000000000")

	// Assert
	assert.Error(t, err)
	assert.NotNil(t, err)
	assert.Equal(t, errs.NotFound, err)

	// Verifique se todas as expectativas do mock foram cumpridas
	mockCardRepo.AssertExpectations(t)
	mockCardRepo.AssertNotCalled(t, "Delete")
}
