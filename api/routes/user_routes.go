package routes

import (
	"github.com/Uallessonivo/go_card_manager/api/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app fiber.Router) {
	httpHandler := &handler.UserHandler{}
	app.Post("/user", httpHandler.CreateUser)
}
