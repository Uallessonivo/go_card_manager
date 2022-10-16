package main

import (
	"github.com/Uallessonivo/go_card_manager/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Listen(":9090")
}
