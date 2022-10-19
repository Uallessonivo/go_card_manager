package main

import (
	"log"

	"github.com/Uallessonivo/go_card_manager/api/routes"
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

func main() {
	database.ConnectDB()
	app := fiber.New()

	api := app.Group("/api")

	routes.UserRoutes(api)

	app.Listen(":9090")
}
