package repository

import (
	"fmt"

	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

func NewUserRepo(Conn *gorm.DB) model.UserRepository {
	return &userRepository{Conn}
}

func (u *userRepository) Create(input *model.User) error {
	err := u.Conn.Create(&input)

	if err != nil {
		return err.Error
	}

	return nil
}

func (u *userRepository) Delete(id string) error {
	err := u.Conn.Delete(&id)

	if err != nil {
		return err.Error
	}

	return nil
}

func (u *userRepository) GetByID(id string) (*model.User, error) {
	var userModel model.User

	u.Conn.First(&userModel, "id = ?", id)

	if userModel.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &userModel, nil
}

func (u *userRepository) Update(input *model.User) error {
	err := u.Conn.Save(input)

	if err != nil {
		return err.Error
	}

	return nil
}
