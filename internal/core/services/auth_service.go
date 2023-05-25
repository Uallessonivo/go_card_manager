package services

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	ports2 "github.com/Uallessonivo/go_card_manager/internal/core/ports"
)

type AuthUseCase struct {
	Auth        ports2.AuthService
	UserService ports2.UserService
}

func NewAuthService(u ports2.UserService) ports2.AuthService {
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
