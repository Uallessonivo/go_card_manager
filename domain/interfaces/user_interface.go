package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type UserUseCaseInterface interface {
	Create(name string, email string, password string) (*model.UserResponse, error)
	GetByID(id string) (*model.UserResponse, error)
}

type UserRepositoryInterface interface {
	Create(input *model.User) error
	GetByID(id string) (*model.User, error)
}
