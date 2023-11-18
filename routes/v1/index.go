package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/initializers"
)

func SetupRoutesV1(app *fiber.App, customValidator *common.CustomValidator ,config *initializers.Config) {
	SetupRoutesAuthV1(app, customValidator, config)
}