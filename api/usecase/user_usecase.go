package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type userUseCase struct {
	userRepository model.UserRepository
}

func NewUserUseCase(u model.UserRepository) model.UserUseCase {
	return &userUseCase{
		userRepository: u,
	}
}

func (u *userUseCase) Create(input *model.User) (*model.User, error) {
	data := model.User{
		ID:       "sdasdfasdasdas",
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	err := u.userRepository.Create(&data)

	if err != nil {
		return nil, err
	}
	return &data, nil
}
