package services_test

import (
	"testing"

	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/services"
	"github.com/stretchr/testify/assert"
)

func TestSaveData_Success(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockCardService := new(MockCardService)

	cardRequests := []*models.CardRequest{
		{Serial: "000000000000000", Owner: "John", Type: "DESPESAS_MATRIZ"},
		{Serial: "000000000000000", Owner: "Doe", Type: "INCENTIVO_CAP1"},
	}

	mockCardService.On("CreateCard", cardRequests[0]).Return(&models.Card{}, nil)
	mockCardService.On("CreateCard", cardRequests[1]).Return(&models.Card{}, nil)

	service := services.NewFileService(mockEmployeeRepo, mockCardRepo, mockCardService)

	// Act
	result, err := service.SaveData(cardRequests)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "All cards have been saved in the postgres", result.Message)
	assert.Empty(t, result.FailedCards)

	mockCardService.AssertExpectations(t)
}

func TestSaveData_Failure(t *testing.T) {
	// Arrange
	mockCardRepo := new(MockCardRepository)
	mockEmployeeRepo := new(MockEmployeeRepository)
	mockCardService := new(MockCardService)

	cardRequests := []*models.CardRequest{
		{Serial: "000000000000000", Owner: "John", Type: "DESPESAS_MATRIZ"},
		{Serial: "000000000000000", Owner: "Doe", Type: "INCENTIVO_CAP1"},
	}

	mockCardService.On("CreateCard", cardRequests[0]).Return(nil, assert.AnError)
	mockCardService.On("CreateCard", cardRequests[1]).Return(&models.Card{}, nil)

	service := services.NewFileService(mockEmployeeRepo, mockCardRepo, mockCardService)

	// Act
	result, err := service.SaveData(cardRequests)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Some cards are not inserted into the postgres", result.Message)
	assert.Len(t, result.FailedCards, 1)
	assert.Equal(t, cardRequests[0], result.FailedCards[0])

	mockCardService.AssertExpectations(t)
}
