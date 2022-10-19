package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type UserUseCase struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserUseCase(u interfaces.UserRepositoryInterface) interfaces.UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: u,
	}
}

func (u *UserUseCase) Create(name string, email string, password string) (*model.User, error) {
	newUser, err := model.NewUser(name, email, password)

	if err != nil {
		return nil, err
	}

	er := u.UserRepository.Create(newUser)

	if er != nil {
		return nil, err
	}

	return newUser, nil
}
