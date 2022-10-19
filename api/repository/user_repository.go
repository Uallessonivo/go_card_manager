package repository

import (
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (u *UserRepositoryDb) Create(input *model.User) error {
	err := u.Db.Create(input).Error

	if err != nil {
		return err
	}

	return nil
}
