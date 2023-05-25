package repositories

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) ports.UserRepository {
	return &UserRepositoryDb{Db}
}

func (u *UserRepositoryDb) Create(input *models.User) error {
	err := u.Db.Create(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) GetByID(id string) (*models.User, error) {
	var user models.User

	err := u.Db.First(&user, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.Db.First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepositoryDb) Update(input *models.User) error {
	err := u.Db.Model(models.User{}).Where("id = ?", input.ID).UpdateColumns(input).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryDb) Delete(id string) error {
	err := u.Db.Delete(&models.User{}, "id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}
