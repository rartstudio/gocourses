package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/services"
)

type authController struct {
	authService *services.AuthService
}

type AuthController interface {
	Register(ctx *fiber.Ctx) error 
}

func NewAuthController(authService *services.AuthService) AuthController {
	return &authController{authService: authService}
}

func (c authController) Register(ctx *fiber.Ctx) error {
	body := new(models.RegisterRequest)
	err := ctx.BodyParser(&body)

	jwtToken, err := c.authService.Register(body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Gagal Registrasi",
		})		
	}

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: fiber.Map{
			"access_token": jwtToken,
		},
		Success: true,
		Message: "Sukses registrasi",
	})
}