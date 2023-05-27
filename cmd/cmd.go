package cmd

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/repositories"
	"github.com/Uallessonivo/go_card_manager/internal/core/services"
	"github.com/Uallessonivo/go_card_manager/internal/infra/postgres"
	"log"

	"github.com/Uallessonivo/go_card_manager/application/routes"
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
	postgres.ConnectDB()

	app := fiber.New()

	// USERS REPO
	uRepo := repositories.NewUserRepository(postgres.DB.Db)
	// CARDS REPO
	cRepo := repositories.NewCardRepository(postgres.DB.Db)
	// EMPLOYEE REPO
	eRepo := repositories.NewEmployeeRepository(postgres.DB.Db)

	// USERS USE CASE
	uCase := services.NewUserService(uRepo)
	// CARDS USE CASE
	cCase := services.NewCardService(cRepo, eRepo)
	// EMPLOYEES USE CASE
	eCase := services.NewEmployeeService(eRepo, cRepo)
	// FILE USE CASE
	fCase := services.NewFileService(eRepo, cRepo, cCase)
	// AUTH USE CASE
	aCase := services.NewAuthService(uRepo)

	// ROUTES
	routes.UserRoutes(app, uCase)
	routes.CardRoutes(app, cCase)
	routes.EmployeeRoutes(app, eCase)
	routes.FileRoutes(app, fCase)
	routes.AuthRoute(app, aCase)

	err := app.Listen(":9090")
	if err != nil {
		return
	}
}
