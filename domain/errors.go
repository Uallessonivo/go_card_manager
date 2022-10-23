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
	ErrInvalidPassword     = errors.New("password must be alphanumeric and must consists of at least 6 characters and not more than 15 characters")
)
