package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/rartstudio/gocourses/routes/v1"
)

func SetupRoutes(app *fiber.App) {
	// health check route
	app.Get("/health-checker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"app": fiber.Map{
				"status": true,
				"message": "Application is up",
			},
		})
	})

	v1.SetupRoutesV1(app)
}