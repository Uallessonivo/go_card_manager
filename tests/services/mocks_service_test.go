package services_test

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
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
