package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestListCardsHandler_Success(t *testing.T) {
	// Arrange
	app := fiber.New()

	// Mock Service
	mockService := new(MockCardService)

	// Mock ListAllCards method
	mockService.On("ListAllCards").Return([]*models.CardResponse{
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
	}, nil)

	// Create handler to register route
	handler := &handlers.CardHandler{
		CardService: mockService,
	}

	app.Get("/cards", handler.ListCards)

	// Http request using httptest
	req := httptest.NewRequest("GET", "/cards", nil)

	// Act: Send request to the handler
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	// Verify response body
	var results []models.CardResponse
	err = json.NewDecoder(resp.Body).Decode(&results)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, "1", results[0].ID)
	assert.Equal(t, "DESPESAS_MATRIZ", results[0].Type)
	assert.Equal(t, "00000000000", results[0].Owner)
	assert.Equal(t, "000000000000000", results[0].Serial)
	assert.Equal(t, "2", results[1].ID)
	assert.Equal(t, "DESPESAS_FILIAL", results[1].Type)
	assert.Equal(t, "11111111111", results[1].Owner)
	assert.Equal(t, "111111111111111", results[1].Serial)
	// Verify all expectations
	mockService.AssertExpectations(t)
}

func TestListCardsHandler_Error(t *testing.T) {
	// Arrange
	app := fiber.New()

	// Mock Service
	mockService := new(MockCardService)

	// Mock ListAllCards method to return an error
	mockService.On("ListAllCards").Return([]*models.CardResponse(nil), errors.New("service error"))

	// Create handler to register route
	handler := &handlers.CardHandler{
		CardService: mockService,
	}

	app.Get("/cards", handler.ListCards)

	// Http request using httptest
	req := httptest.NewRequest("GET", "/cards", nil)

	// Act: Send request to the handler
	resp, err := app.Test(req, -1)
	assert.NoError(t, err)

	// Assert
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	// Verify response body
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "service error", result["error"])
	// Verify all expectations
	mockService.AssertExpectations(t)
}
