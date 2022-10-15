package domain

import "context"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	Create(ctx context.Context, input *User) error
	Update(ctx context.Context, input *User) error
	Delete(ctx context.Context, id string) error
}
