package interfaces

import "github.com/Uallessonivo/go_card_manager/domain/model"

type UserRepositoryInterface interface {
	Create(input *model.User) error
}
