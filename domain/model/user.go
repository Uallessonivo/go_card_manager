package model

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	GetByID(id string) (*User, error)
	Create(input *User) error
	Update(input *User) error
	Delete(id string) error
}
