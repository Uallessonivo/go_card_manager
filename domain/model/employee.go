package model

type Employee struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Cpf   string  `json:"cpf"`
	Cards []*Card `gorm:"foreignKey:Owner"`
}

type EmployeeRequest struct {
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}
