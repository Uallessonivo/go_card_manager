package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/entities"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) interfaces.UserRepositoryInterface {
	return &UserRepositoryDb{Db}
}

func (u *UserRepositoryDb) Create(input *entities.User) error {
	err := u.Db.Create(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) GetByID(id string) (*entities.User, error) {
	var user entities.User

	err := u.Db.First(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) GetByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := u.Db.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) Update(input *entities.User) error {
	err := u.Db.Model(entities.User{}).Where("id = ?", input.ID).UpdateColumns(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) Delete(id string) error {
	err := u.Db.Delete(&entities.User{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
