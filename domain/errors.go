package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("user not found")
	ErrInvalid             = errors.New("something went wrong")
	ErrConflict            = errors.New("this item already exists")
	ErrUserExists          = errors.New("this user already exists")
	ErrBadParamInput       = errors.New("given param is not valid")
)
