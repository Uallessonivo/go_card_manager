package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/entities"

type AuthUseCaseInterface interface {
	Login(input *entities.LoginRequest) (*entities.LoginResponse, error)
	GenerateJWT() error
	ValidateJWT() error
}
