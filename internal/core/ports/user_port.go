package ports

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
)

type UserService interface {
	CreateUser(input *models.UserRequest) (*models.UserResponse, error)
	GetUserByID(id string) (*models.UserResponse, error)
	GetUserByEmail(email string) (*models.UserResponse, error)
	CheckUserPass(email string, password string) error
	UpdateUser(id string, input *models.UserRequest) (*models.UserResponse, error)
	DeleteUser(id string) error
}

type UserRepository interface {
	Create(input *models.User) error
	GetByID(id string) (*models.User, error)
	Update(input *models.User) error
	Delete(id string) error
	GetByEmail(email string) (*models.User, error)
}
