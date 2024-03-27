package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/services"
	"github.com/rartstudio/gocourses/utils"
)

type authController struct {
	authService *services.AuthService
	userService *services.UserService
}

type AuthController interface {
	Register(ctx *fiber.Ctx) error 
	Login(ctx *fiber.Ctx) error
	Verify(ctx *fiber.Ctx) error
}

func NewAuthController(authService *services.AuthService, userService *services.UserService) AuthController {
	return &authController{authService: authService, userService: userService}
}

func (c authController) Register(ctx *fiber.Ctx) error {
	body := new(models.RegisterRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	jwtToken, err := c.authService.Register(body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: err.Error(),
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

func (c authController) Login(ctx *fiber.Ctx) error {
	body := new(models.LoginRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	model, err := c.userService.GetByEmail(body.Email)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Akun tidak ditemukan",
		})		
	}

	jwtToken, err := c.authService.Login(body, model)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: err.Error(),
		})		
	}

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: fiber.Map{
			"access_token": jwtToken,
		},
		Success: true,
		Message: "Sukses login",
	})
}

func (c authController) Verify(ctx *fiber.Ctx) error {
	body := new(models.VerifyAccountRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	jwtClaims := ctx.Locals("user").(utils.UserCredential)

	model, err := c.authService.VerifyAccount(body, jwtClaims.ID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Gagal verifikasi pengguna",
		})
	}

	response := models.FilterUserResponseV1(model)

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: response,
		Success: true,
		Message: "Sukses verifikasi pengguna",
	})
}
