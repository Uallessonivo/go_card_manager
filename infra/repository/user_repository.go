package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepositoryDb{Db}
}

func (u *UserRepositoryDb) Create(input *model.User) error {
	err := u.Db.Create(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) GetByID(id string) (*model.User, error) {
	var user model.User

	err := u.Db.First(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) GetByEmail(email string) (*model.User, error) {
	var user model.User

	err := u.Db.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) Update(input *model.User) error {
	err := u.Db.Model(model.User{}).Where("id = ?", input.ID).UpdateColumns(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) Delete(id string) error {
	err := u.Db.Delete(&model.User{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
