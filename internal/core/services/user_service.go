package services

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/models"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
)

type UserUseCase struct {
	UserRepository ports.UserRepository
}

func NewUserService(u ports.UserRepository) ports.UserService {
	return &UserUseCase{
		UserRepository: u,
	}
}

func (u UserUseCase) CreateUser(input *models.UserRequest) (*models.UserResponse, error) {
	newUser, err := models.MakeUser(input)
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

	return &models.UserResponse{
		ID:    newUser.ID,
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (u UserUseCase) GetUserByID(id string) (*models.UserResponse, error) {
	userFound, err := u.UserRepository.GetByID(id)

	if err != nil {
		return nil, errors.NotFound
	}

	response := models.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u UserUseCase) GetUserByEmail(email string) (*models.UserResponse, error) {
	userFound, err := u.UserRepository.GetByEmail(email)

	if err != nil {
		return nil, errors.NotFound
	}

	response := models.UserResponse{
		ID:    userFound.ID,
		Name:  userFound.Name,
		Email: userFound.Email,
	}

	return &response, nil
}

func (u UserUseCase) CheckUserPass(email string, password string) error {
	user, err := u.UserRepository.GetByEmail(email)
	if err != nil {
		return errors.NotFound
	}

	hashedPass, err := models.HashPassword(password)
	if err != nil {
		return err
	}

	if user.Password != hashedPass {
		return errors.InvalidLogin
	}

	return nil
}

func (u UserUseCase) UpdateUser(id string, input *models.UserRequest) (*models.UserResponse, error) {
	_, err := u.UserRepository.GetByID(id)
	if err != nil {
		return nil, errors.NotFound
	}

	if err := models.ValidateUser(input); err != nil {
		return nil, err
	}

	if err := u.UserRepository.Update(&models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:    id,
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

func (u UserUseCase) DeleteUser(id string) error {
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
