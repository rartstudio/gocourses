package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/services"
	"github.com/rartstudio/gocourses/utils"
)

type userController struct {
	userService *services.UserService
}

type UserController interface {
	User(ctx *fiber.Ctx) error
}

func NewUserController(userService *services.UserService) UserController {
	return &userController{userService: userService}
}

func (c userController) User(ctx *fiber.Ctx) error {
	jwtClaims := ctx.Locals("user").(utils.UserCredential)

	model, err := c.userService.GetByUuid(jwtClaims.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Pengguna tidak ditemukan",
		})
	}

	response := models.FilterUserResponseV1(model)

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: response,
		Success: true,
		Message: "Sukses verifikasi pengguna",
	})
}