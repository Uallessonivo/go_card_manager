package routes

import (
	"github.com/Uallessonivo/go_card_manager/application/handler"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, us interfaces.UserUseCaseInterface) {
	httpHandler := &handler.UserHandler{
		UseCase: us,
	}
	app.Post("/user/create", httpHandler.CreateUser)
	app.Get("/user/find/:id", httpHandler.GetUserByID)
	app.Get("/user/find/:email", httpHandler.GetUserByEmail)
	app.Delete("/user/delete/:id", httpHandler.DeleteUser)
	app.Put("/user/update/:id", httpHandler.UpdateUser)
}
