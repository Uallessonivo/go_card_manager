package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
)

type AuthUseCase struct {
	Auth        interfaces.AuthUseCaseInterface
	UserUseCase interfaces.UserUseCaseInterface
}

func NewAuthUseCase(u interfaces.UserUseCaseInterface) interfaces.AuthUseCaseInterface {
	return &AuthUseCase{
		UserUseCase: u,
	}
}

// Login implements interfaces.AuthUseCaseInterface
func (a AuthUseCase) Login(input *entities.LoginRequest) (*entities.LoginResponse, error) {
	panic("unimplemented")
}

// GenerateToken implements interfaces.AuthUseCaseInterface
func (a AuthUseCase) GenerateJWT() error {
	panic("unimplemented")
}

// ValidateJWT implements interfaces.AuthUseCaseInterface
func (a AuthUseCase) ValidateJWT() error {
	panic("unimplemented")
}
