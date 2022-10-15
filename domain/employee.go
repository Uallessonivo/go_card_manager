package domain

import "context"

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
}

type EmployeeRepository interface {
	List() ([]Employee, error)
	GetByID(ctx context.Context, input string) (Employee, error)
	GetByCpf(ctx context.Context, input string) (Employee, error)
	GetByName(ctx context.Context, input string) (Employee, error)
	Create(ctx context.Context, input *Employee) error
	Update(ctx context.Context, input *Employee) error
	Delete(ctx context.Context, id string) error
}
