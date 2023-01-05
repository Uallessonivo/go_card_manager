package routes

import (
	"github.com/Uallessonivo/go_card_manager/application/handler"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App, us interfaces.EmployeeUseCaseInterface) {
	httpHandler := &handler.EmployeeHandler{
		UseCase: us,
	}

	app.Post("/employee/create", httpHandler.CreateEmployee)
}
