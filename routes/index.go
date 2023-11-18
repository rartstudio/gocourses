package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/rartstudio/gocourses/routes/v1"
)

func SetupRoutes(app *fiber.App) {
	v1.SetupRoutesV1(app)
}