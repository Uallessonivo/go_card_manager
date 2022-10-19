package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type UserUseCaseInterface interface {
	Create(name string, email string, password string) (*model.User, error)
}

type UserRepositoryInterface interface {
	Create(input *model.User) error
}
