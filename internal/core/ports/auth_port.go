package ports

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type AuthService interface {
	Login(input *models.LoginRequest) (*models.LoginResponse, error)
	GenerateJWT(id string) (string, error)
	ValidateJWT(tokenString string) error
}
