package usecase

import (
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

func (u *UserUseCase) Create(name string, email string, password string) (*model.UserResponse, error) {
	newUser, err := model.MakeUser(name, email, password)

	if err != nil {
		return nil, err
	}

	userFound, foundErr := u.UserRepository.GetByID(newUser.ID)

	if foundErr != nil || userFound.ID != "" || userFound.Email != "" {
		return nil, foundErr
	}

	er := u.UserRepository.Create(newUser)

	if er != nil {
		return nil, err
	}

	response := model.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}

	return &response, nil
}

func (u *UserUseCase) GetByID(id string) (*model.UserResponse, error) {
	userFound, err := u.UserRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	response := model.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u *UserUseCase) Update(name string, email string, password string) (*model.UserResponse, error) {
	updateUser, err := model.MakeUser(name, email, password)

	if err != nil {
		return nil, err
	}

	er := u.UserRepository.Update(updateUser)

	if er != nil {
		return nil, err
	}

	response := model.UserResponse{
		ID:    updateUser.ID,
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}

	return &response, nil
}

func (u *UserUseCase) Delete(id string) error {
	userFound, err := u.UserRepository.GetByID(id)

	if err != nil {
		return err
	}

	er := u.UserRepository.Delete(userFound.ID)

	if er != nil {
		return er
	}

	return nil
}
