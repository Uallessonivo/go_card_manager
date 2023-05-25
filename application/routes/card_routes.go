package routes

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func CardRoutes(app *fiber.App, us ports.CardService) {
	httpHandler := &handlers.CardHandler{
		CardService: us,
	}

	app.Post("/card/create", httpHandler.CreateCard)
	app.Get("/cards/list", httpHandler.ListCards)
	app.Get("/cards/filter-by-type/:type", httpHandler.ListCardsByType)
	app.Get("/cards/filter-by-owner/:owner", httpHandler.ListCardsByOwner)
	app.Delete("/card/delete/:id", httpHandler.DeleteCard)
}
