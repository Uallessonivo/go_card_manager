package model

import (
	"github.com/Uallessonivo/go_card_manager/domain"
	uuid "github.com/satori/go.uuid"
	"regexp"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func emailIsValid(email string) error {
	_, err := regexp.MatchString(`^[\w-]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	if err != nil {
		return err
	}
	return nil
}

func passwordIsValid(password string) error {
	_, err := regexp.MatchString(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*[a-zA-Z]).{8,}$`, password)
	if err != nil {
		return err
	}
	return nil
}

func MakeUser(name string, email string, password string) (*User, error) {
	err := emailIsValid(email)
	if err != nil {
		return nil, domain.ErrInvalidEmail
	}

	errr := passwordIsValid(password)
	if errr != nil {
		return nil, domain.ErrInvalidPassword
	}

	newUser := User{
		ID:       uuid.NewV4().String(),
		Name:     name,
		Email:    email,
		Password: password,
	}

	return &newUser, nil
}
