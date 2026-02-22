package main

import (
	"log"

	"manga-downloader/middlewares"
	"manga-downloader/routes"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		StructValidator: middlewares.NewValidator(),
	})

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3001"))
}
