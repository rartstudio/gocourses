package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/initializers"
	v1 "github.com/rartstudio/gocourses/routes/v1"
)

func SetupRoutes(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config) {
	v1.SetupRoutesV1(app, customValidator, config)
}