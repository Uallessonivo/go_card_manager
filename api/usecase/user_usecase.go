package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type UserUseCase struct {
	UserRepository model.UserRepository
}

func (u *UserUseCase) Create(name string, email string, password string) (*model.User, error) {
	data := model.User{
		ID:       "sdasdfasdasdas",
		Name:     name,
		Email:    email,
		Password: password,
	}
	err := u.UserRepository.Create(&data)

	if err != nil {
		return nil, err
	}
	return &data, nil
}
