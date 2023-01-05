package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type UserUseCaseInterface interface {
	Create(input *model.UserRequest) (*model.UserResponse, error)
	GetByID(id string) (*model.UserResponse, error)
	GetByEmail(email string) (*model.UserResponse, error)
	Update(id string, input *model.UserRequest) (*model.UserResponse, error)
	Delete(id string) error
}

type UserRepositoryInterface interface {
	Create(input *model.User) error
	GetByID(id string) (*model.User, error)
	Update(input *model.User) error
	Delete(id string) error
	GetByEmail(email string) (*model.User, error)
}
