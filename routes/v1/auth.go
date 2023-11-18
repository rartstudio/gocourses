package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
)

func SetupRoutesAuthV1(app *fiber.App) {
	apiV1 := app.Group("/api/v1/auth")
	apiV1.Post("/register", func(c *fiber.Ctx) error {
		return c.JSON(common.SuccessHandlerResp{
			Success: true,
			Message: "Sukses registrasi",
			Data: nil,
		})
	})
	apiV1.Post("/login",func(c *fiber.Ctx) error {
		return c.JSON(common.SuccessHandlerResp{
			Success: true,
			Message: "Sukses login",
			Data: nil,
		})
	})
	apiV1.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(common.SuccessHandlerResp{
			Data: nil,
			Success: true,
			Message: "Sukses mendapatkan profil user",
		})
	})
}