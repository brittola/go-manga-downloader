package controllers

import "github.com/gofiber/fiber/v3"

func Health(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API is running",
		"status":  "success",
	})
}
