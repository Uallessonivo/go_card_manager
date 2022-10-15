package main

import (
	"github.com/Uallessonivo/go_card_manager/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong")
	})

	app.Listen(":9090")
}
