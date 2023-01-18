package usecase

import "github.com/Uallessonivo/go_card_manager/domain/interfaces"

type FileUseCase struct {
	EmployeeRepository interfaces.EmployeeRepositoryInterface
	CardRepository     interfaces.CardRepositoryInterface
}

func NewFileUseCase(e interfaces.EmployeeRepositoryInterface, c interfaces.CardRepositoryInterface) interfaces.FileUseCaseInterface {
	return &FileUseCase{
		EmployeeRepository: e,
		CardRepository:     c,
	}
}

func (f FileUseCase) SaveData() error {
	//TODO implement me
	panic("implement me")
}

func (f FileUseCase) GenerateCardsReport() error {
	//TODO implement me
	panic("implement me")
}

func (f FileUseCase) GenerateEmployeesReport() error {
	//TODO implement me
	panic("implement me")
}
