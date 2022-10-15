package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("not found")
	ErrConflict            = errors.New("this item already exists")
	ErrBadParamInput       = errors.New("given param is not a valid")
)
