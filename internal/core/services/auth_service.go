package services

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	ports "github.com/Uallessonivo/go_card_manager/internal/core/ports"
)

type AuthUseCase struct {
	Auth        ports.AuthService
	UserService ports.UserService
}

func NewAuthService(u ports.UserService) ports.AuthService {
	return &AuthUseCase{
		UserService: u,
	}
}

// Login implements ports.AuthService
func (a AuthUseCase) Login(input *models.LoginRequest) (*models.LoginResponse, error) {
	panic("unimplemented")
}

// GenerateJWT implements ports.AuthService
func (a AuthUseCase) GenerateJWT() error {
	panic("unimplemented")
}

// ValidateJWT implements ports.AuthService
func (a AuthUseCase) ValidateJWT() error {
	panic("unimplemented")
}
