package cmd

import (
	"log"

	"github.com/Uallessonivo/go_card_manager/api/repository"
	"github.com/Uallessonivo/go_card_manager/api/routes"
	"github.com/Uallessonivo/go_card_manager/api/usecase"
	"github.com/Uallessonivo/go_card_manager/database"
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

	uRepo := repository.NewUserRepository(database.DB.Db)
	uCase := usecase.NewUserUseCase(uRepo)

	routes.UserRoutes(app, uCase)

	app.Listen(":9090")
}
