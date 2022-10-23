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

func emailIsValid(email string) bool {
	matched, _ := regexp.MatchString(`^[\w-]+@([\w-]+\.)+[\w-]{2,4}$`, email)
	return matched
}

func passwordIsValid(password string) bool {
	matched, _ := regexp.MatchString(`^([a-zA-Z0-9@*#]{6,15})$`, password)
	return matched
}

func MakeUser(id string, name string, email string, password string) (*User, error) {
	emailIsValid := emailIsValid(email)
	if emailIsValid == false {
		return nil, domain.ErrInvalidEmail
	}

	passwordIsValid := passwordIsValid(password)
	if passwordIsValid == false {
		return nil, domain.ErrInvalidPassword
	}

	if id == "" {
		id = uuid.NewV4().String()
	}

	newUser := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: password,
	}

	return &newUser, nil
}
