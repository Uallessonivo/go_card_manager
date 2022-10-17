package model

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

type EmployeeRepository interface {
	List() ([]*Employee, error)
	GetByID(input string) (*Employee, error)
	GetByCpf(input string) (*Employee, error)
	GetByName(input string) (*Employee, error)
	Create(input *Employee) error
	Update(input *Employee) error
	Delete(id string) error
}
