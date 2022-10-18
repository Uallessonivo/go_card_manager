package repository

import (
	"fmt"

	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(Conn *gorm.DB) model.UserRepository {
	return &UserRepository{Conn}
}

func (u *UserRepository) Create(input *model.User) error {
	err := u.Conn.Create(&input)

	if err != nil {
		return err.Error
	}

	return nil
}

func (u *UserRepository) Delete(id string) error {
	err := u.Conn.Delete(&id)

	if err != nil {
		return err.Error
	}

	return nil
}

func (u *UserRepository) GetByID(id string) (*model.User, error) {
	var userModel model.User

	u.Conn.First(&userModel, "id = ?", id)

	if userModel.ID == "" {
		return nil, fmt.Errorf("user not found")
	}

	return &userModel, nil
}

func (u *UserRepository) Update(input *model.User) error {
	err := u.Conn.Save(input)

	if err != nil {
		return err.Error
	}

	return nil
}
