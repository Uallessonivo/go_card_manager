package usecase

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/Uallessonivo/go_card_manager/domain/model"
	"os"
)

type UserUseCase struct {
	UserRepository interfaces.UserRepositoryInterface
}

func NewUserUseCase(u interfaces.UserRepositoryInterface) interfaces.UserUseCaseInterface {
	return &UserUseCase{
		UserRepository: u,
	}
}

func (u *UserUseCase) Create(name string, email string, password string, secretKey string) (*model.UserResponse, error) {
	newUser, err := model.MakeUser("", name, email, password)
	if err != nil {
		return nil, err
	}

	userExists, _ := u.UserRepository.GetByEmail(newUser.Email)
	if userExists != nil {
		return nil, errors.UserExists
	}

	if secretKey != os.Getenv("SECRET_KEY") {
		return nil, errors.InvalidParams
	}

	er := u.UserRepository.Create(newUser)
	if er != nil {
		return nil, er
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
		return nil, errors.UserNotFound
	}

	response := model.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u *UserUseCase) GetByEmail(email string) (*model.UserResponse, error) {
	userFound, err := u.UserRepository.GetByEmail(email)

	if err != nil {
		return nil, errors.UserNotFound
	}

	response := model.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u *UserUseCase) Update(id string, name string, email string, password string) (*model.UserResponse, error) {
	_, errr := u.UserRepository.GetByID(id)
	if errr != nil {
		return nil, errors.UserNotFound
	}

	updateUser, updateUserErr := model.MakeUser(id, name, email, password)
	if updateUserErr != nil {
		return nil, updateUserErr
	}

	er := u.UserRepository.Update(updateUser)
	if er != nil {
		return nil, er
	}

	response := model.UserResponse{
		ID:    updateUser.ID,
		Name:  updateUser.Name,
		Email: updateUser.Email,
	}

	return &response, nil
}

func (u *UserUseCase) Delete(id string) error {
	_, err := u.UserRepository.GetByID(id)
	if err != nil {
		return errors.UserNotFound
	}

	er := u.UserRepository.Delete(id)
	if er != nil {
		return er
	}

	return nil
}
