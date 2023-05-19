package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
)

type UserUseCase struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserUseCase(u interfaces.UserRepositoryInterface) interfaces.UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: u,
	}
}

func (u *UserUseCase) CreateUser(input *model.UserRequest) (*model.UserResponse, error) {
	newUser, err := model.MakeUser(input)
	if err != nil {
		return nil, err
	}

	userExists, _ := u.UserRepository.GetByEmail(newUser.Email)
	if userExists != nil {
		return nil, errors.AlreadyExists
	}

	er := u.UserRepository.Create(newUser)
	if er != nil {
		return nil, er
	}

	return &model.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (u *UserUseCase) GetUserByID(id string) (*model.UserResponse, error) {
	userFound, err := u.UserRepository.GetByID(id)

	if err != nil {
		return nil, errors.NotFound
	}

	response := model.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u *UserUseCase) GetUserByEmail(email string) (*model.UserResponse, error) {
	userFound, err := u.UserRepository.GetByEmail(email)

	if err != nil {
		return nil, errors.NotFound
	}

	response := model.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u *UserUseCase) UpdateUser(id string, input *model.UserRequest) (*model.UserResponse, error) {
	_, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, errors.NotFound
	}

	if err := model.ValidateUser(input); err != nil {
		return nil, err
	}

	if err := u.UserRepository.Update(&model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	return &model.UserResponse{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

func (u *UserUseCase) DeleteUser(id string) error {
	_, err := u.UserRepository.GetByID(id)
	if err != nil {
		return errors.NotFound
	}

	er := u.UserRepository.Delete(id)
	if er != nil {
		return er
	}

	return nil
}
