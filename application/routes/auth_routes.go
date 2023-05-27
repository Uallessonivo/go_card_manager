package routes

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, us ports.AuthService) {
	httpHandler := &handlers.AuthHandler{
		AuthService: us,
	}

	app.Post("/authenticate", httpHandler.Authenticate)
}
