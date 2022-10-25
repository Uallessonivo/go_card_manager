package domain

import "errors"

var (
	ErrUserNotFound    = errors.New("user not found")
	ErrUserExists      = errors.New("this user already exists")
	ErrInvalidEmail    = errors.New("email is invalid")
	ErrInvalidPassword = errors.New("password must be alphanumeric and must consists of at least 6 characters and not more than 15 characters")
	ErrInvalidParams = errors.New("given params is invalid")
)
