package routes

import (
	"github.com/Uallessonivo/go_card_manager/application/handler"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

func CardRoutes(app *fiber.App, us interfaces.CardUseCaseInterface) {
	httpHandler := &handler.CardHandler{
		UseCase: us,
	}

	app.Post("/card/create", httpHandler.CreateCard)
	app.Get("/cards/list", httpHandler.ListCards)
	app.Get("/cards/filter-by-type/:type", httpHandler.ListCardsByType)
	app.Get("/cards/filter-by-owner/:owner", httpHandler.ListCardsByOwner)
	app.Delete("/card/delete/:id", httpHandler.DeleteCard)
}
