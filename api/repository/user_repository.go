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
	err := u.Db.First(&model.User{}, id).Error

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (u *UserRepositoryDb) Update(input *model.User) error {
	err := u.Db.Save(&model.User{}).Error

	if err != nil {
		return err
	}

	return nil

}

func (u *UserRepositoryDb) Delete(id string) error {
	err := u.Db.Delete(&model.User{}, id).Error

	if err != nil {
		return err
	}

	return nil

}
