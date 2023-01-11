package cmd

import (
	"github.com/Uallessonivo/go_card_manager/infra/database"
	repository "github.com/Uallessonivo/go_card_manager/infra/repository"
	"log"

	"github.com/Uallessonivo/go_card_manager/application/routes"
	"github.com/Uallessonivo/go_card_manager/application/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading environment file")
	}
}

func Execute() {
	database.ConnectDB()

	app := fiber.New()

	// USERS REPO
	uRepo := repository.NewUserRepository(database.DB.Db)
	// CARDS REPO
	cRepo := repository.NewCardRepository(database.DB.Db)
	// EMPLOYEE REPO
	eRepo := repository.NewEmployeeRepository(database.DB.Db)

	// USERS USE CASE
	uCase := usecase.NewUserUseCase(uRepo)
	// CARDS USE CASE
	cCase := usecase.NewCardUseCase(cRepo, eRepo)
	// EMPLOYEES USE CASE
	eCase := usecase.NewEmployeeUseCase(eRepo, cRepo)

	routes.UserRoutes(app, uCase)
	routes.CardRoutes(app, cCase)
	routes.EmployeeRoutes(app, eCase)

	err := app.Listen(":9090")
	if err != nil {
		return
	}
}
