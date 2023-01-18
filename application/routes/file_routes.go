package routes

import (
	"github.com/Uallessonivo/go_card_manager/application/handler"
	"github.com/Uallessonivo/go_card_manager/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

func FileRoutes(app *fiber.App, us interfaces.FileUseCaseInterface) {
	fileHandler := &handler.FileHandler{
		UseCase: us,
	}

	app.Post("/file/upload", fileHandler.UploadFile)
	app.Get("/card-report/download", fileHandler.DownloadCardReport)
	app.Get("/employee-report/download", fileHandler.DownloadEmployeeReport)
}
