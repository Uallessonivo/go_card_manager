package model

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

func NewUser(name string, email string, password string) (*User, error) {
	newUser := User{
		ID:       "sadasdasdasd",
		Name:     name,
		Email:    email,
		Password: password,
	}

	return &newUser, nil
}
