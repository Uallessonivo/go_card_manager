package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/entities"

type UserUseCaseInterface interface {
	CreateUser(input *entities.UserRequest) (*entities.UserResponse, error)
	GetUserByID(id string) (*entities.UserResponse, error)
	GetUserByEmail(email string) (*entities.UserResponse, error)
	CheckUserPass(email string, password string) error
	UpdateUser(id string, input *entities.UserRequest) (*entities.UserResponse, error)
	DeleteUser(id string) error
}

type UserRepositoryInterface interface {
	Create(input *entities.User) error
	GetByID(id string) (*entities.User, error)
	Update(input *entities.User) error
	Delete(id string) error
	GetByEmail(email string) (*entities.User, error)
}
