package models

import (
	"github.com/Uallessonivo/go_card_manager/internal/core/domain/errors"
	"os"
	"regexp"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func ComparePasswords(hashedPassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func ValidateUser(input *UserRequest) error {
	if !emailIsValid(input.Email) {
		return errors.InvalidEmail
	}

	if !passwordIsValid(input.Password) {
		return errors.InvalidPassword
	}

	return nil
}

func MakeUser(input *UserRequest) (*User, error) {
	if input.SecretKey != os.Getenv("SECRET_KEY") {
		return nil, errors.InvalidParams
	}

	if err := ValidateUser(input); err != nil {
		return nil, err
	}

	passwordHash, err := HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.NewV4().String(),
		Name:     input.Name,
		Email:    input.Email,
		Password: passwordHash,
	}, nil
}
