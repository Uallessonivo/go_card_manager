package routes

import (
	"github.com/Uallessonivo/go_card_manager/internal/adapters/handlers"
	"github.com/Uallessonivo/go_card_manager/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

func FileRoutes(app *fiber.App, us ports.FileService) {
	fileHandler := &handlers.FileHandler{
		FileService: us,
	}

	app.Post("/file/upload", fileHandler.UploadFile)
	app.Get("/card-report/download", fileHandler.DownloadCardReport)
	app.Get("/employee-report/download", fileHandler.DownloadEmployeeReport)
}
