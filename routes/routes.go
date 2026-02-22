package routes

import (
	"manga-downloader/controllers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/health", controllers.Health)

	app.Post("/download", controllers.DownloadManga)
}
