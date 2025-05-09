package router

import (
	"github.com/gofiber/fiber/v2"

	"go-fiber-app/service"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/items", service.CreateItem)
	// app.Get("/hello", service.GenerateExcel)
	// app.Get("/registers", service.GetRegisters)
	// app.Get("/export-registers", service.ExportRegistersToExcel)
}
