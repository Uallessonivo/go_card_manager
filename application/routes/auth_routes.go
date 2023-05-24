package routes

import (
	"github.com/Uallessonivo/go_card_manager/application/handler"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(app *fiber.App, us interfaces.AuthUseCaseInterface) {
	httpHandler := &handler.AuthHandler{
		AuthUseCase: us,
	}

	app.Post("/authenticate", httpHandler.Authenticate)
}
