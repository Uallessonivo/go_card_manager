package cmd

import (
	"github.com/Uallessonivo/go_card_manager/infra/database"
	repository2 "github.com/Uallessonivo/go_card_manager/infra/repository"
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

	// USERS
	uRepo := repository2.NewUserRepository(database.DB.Db)
	uCase := usecase.NewUserUseCase(uRepo)
	// CARDS
	cRepo := repository2.NewCardRepository(database.DB.Db)
	cCase := usecase.NewCardUseCase(cRepo)

	routes.UserRoutes(app, uCase)
	routes.CardRoutes(app, cCase)

	err := app.Listen(":9090")
	if err != nil {
		return
	}
}
