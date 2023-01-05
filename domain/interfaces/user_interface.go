package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type UserUseCaseInterface interface {
	CreateUser(input *model.UserRequest) (*model.UserResponse, error)
	GetUserByID(id string) (*model.UserResponse, error)
	GetUserByEmail(email string) (*model.UserResponse, error)
	UpdateUser(id string, input *model.UserRequest) (*model.UserResponse, error)
	DeleteUser(id string) error
}

type UserRepositoryInterface interface {
	Create(input *model.User) error
	GetByID(id string) (*model.User, error)
	Update(input *model.User) error
	Delete(id string) error
	GetByEmail(email string) (*model.User, error)
}
