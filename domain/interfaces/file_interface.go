package interfaces

// TODO

type FileUseCaseInterface interface {
	SaveData() error
	GenerateCardsReport() error
	GenerateEmployeesReport() error
}
