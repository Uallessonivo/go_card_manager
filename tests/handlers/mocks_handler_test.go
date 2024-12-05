package handlers_test

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/stretchr/testify/mock"
)

type MockCardService struct {
	mock.Mock
}

func (m *MockCardService) CreateCard(input *models.CardRequest) (*models.CardResponse, error) {
	args := m.Called(input)
	return args.Get(0).(*models.CardResponse), args.Error(1)
}

func (m *MockCardService) DeleteCard(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCardService) ValidateCard(input string) (*models.Employee, error) {
	args := m.Called(input)
	return args.Get(0).(*models.Employee), args.Error(1)
}

func (m *MockCardService) ListAllCards() ([]*models.CardResponse, error) {
	args := m.Called()
	return args.Get(0).([]*models.CardResponse), args.Error(1)
}

func (m *MockCardService) ListAllCardsByType(input string) ([]*models.CardResponse, error) {
	args := m.Called(input)
	return args.Get(0).([]*models.CardResponse), args.Error(1)
}

func (m *MockCardService) ListAllCardsByOwner(input string) ([]*models.CardResponse, error) {
	args := m.Called(input)
	return args.Get(0).([]*models.CardResponse), args.Error(1)
}
