package model

import (
	"github.com/Uallessonivo/go_card_manager/domain/errors"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strconv"
)

type User struct {
	ID       string `gorm:"primary_key"`
	Name     string
	Email    string
	Password string
}

type UserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	SecretKey string `json:"secret_key"`
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

func hashPassword(password string) (string, error) {
	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func MakeUser(id string, name string, email string, password string) (*User, error) {
	if !emailIsValid(email) {
		return nil, errors.InvalidEmail
	}

	if !passwordIsValid(password) {
		return nil, errors.InvalidPassword
	}

	passwordHash, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	if id == "" {
		id = uuid.NewV4().String()
	}

	newUser := User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: passwordHash,
	}

	return &newUser, nil
}
