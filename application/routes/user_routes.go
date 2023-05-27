package routes

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, us ports.UserService) {
	httpHandler := &handlers.UserHandler{
		UserService: us,
	}
	app.Post("/user/create", httpHandler.CreateUser)
	app.Get("/user/find/id/:id", httpHandler.GetUserByID)
	app.Get("/user/find/email/:email", httpHandler.GetUserByEmail)
	app.Delete("/user/delete/:id", httpHandler.DeleteUser)
	app.Post("/user/update/:id", httpHandler.UpdateUser)
}
