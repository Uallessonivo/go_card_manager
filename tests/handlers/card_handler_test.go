package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
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

func TestCreateCardHandler_Success(t *testing.T) {
	// Arrange
	app := fiber.New()

	// Mock Service
	mockService := new(MockCardService)

	// Mock CreateCard method
	mockService.On("CreateCard", mock.Anything).Return(&models.CardResponse{
		ID:     "1",
		Type:   "DESPESAS_MATRIZ",
		Owner:  "00000000000",
		Serial: "000000000000000",
	}, nil)

	// Create handler to register route
	handler := &handlers.CardHandler{
		CardService: mockService,
	}

	app.Post("/card/create", handler.CreateCard)

	// Request body
	cardRequest := map[string]interface{}{
		"type":   "DESPESAS_MATRIZ",
		"owner":  "00000000000",
		"serial": "000000000000000",
	}
	body, err := json.Marshal(cardRequest)
	assert.NoError(t, err)

	// Http request using httptest
	req := httptest.NewRequest("POST", "/card/create", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	// Act: Send request to the handler
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// Verify request body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "1", result["id"])
	assert.Equal(t, "DESPESAS_MATRIZ", result["type"])
	assert.Equal(t, "000000000000000", result["serial"])
	// Verify all expectations
	mockService.AssertExpectations(t)
}
