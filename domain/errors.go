package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound        = errors.New("user not found")
	ErrInvalid             = errors.New("something went wrong")
	ErrConflict            = errors.New("this item already exists")
	ErrUserExists          = errors.New("this user already exists")
	ErrBadParamInput       = errors.New("given param is not valid")
	ErrInvalidEmail        = errors.New("email is invalid")
	ErrInvalidPassword     = errors.New("passwords must contains at least 8 characters, at least 1 uppercase letter, 1 lowercase letter, and 1 number, can contain special characters")
)
