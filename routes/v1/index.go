package v1

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutesV1(app *fiber.App) {
	SetupRoutesAuthV1(app)
}