package controllers

import (
	"fmt"
	"manga-downloader/models"
	"manga-downloader/services"
	"manga-downloader/utils"

	"github.com/gofiber/fiber/v3"
)

func DownloadManga(c fiber.Ctx) error {
	var request models.DownloadRequest

	if err := c.Bind().Body(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Validation failed",
			"details": utils.FormatValidationErrors(err),
		})
	}

	err, pdfPath := services.DownloadManga(request.Manga, request.Chapter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "error",
		})
	}

	fmt.Println(pdfPath)

	return c.SendFile(pdfPath)
}
