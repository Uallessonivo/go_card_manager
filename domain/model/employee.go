package model

type Employee struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Cpf   string
	Cards []*Card `gorm:"many2many:cards;"`
}

type EmployeeRequest struct {
	Name  string         `json:"name"`
	Cpf   string         `json:"cpf"`
	Cards []*CardRequest `json:"cards"`
}
