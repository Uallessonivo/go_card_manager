package routes

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRoutes(app *fiber.App, us ports.EmployeeService) {
	httpHandler := &handlers.EmployeeHandler{
		EmployeeService: us,
	}

	app.Post("/employee/create", httpHandler.CreateEmployee)
	app.Post("/employee/update/:id", httpHandler.UpdateEmployee)
	app.Get("/employees/list", httpHandler.ListEmployees)
	app.Get("/employee/get", httpHandler.GetEmployee)
	app.Delete("/employee/delete/:id", httpHandler.DeleteEmployee)
}
