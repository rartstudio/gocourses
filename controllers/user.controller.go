package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/clients"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/services"
	"github.com/rartstudio/gocourses/utils"
)

type userController struct {
	userService *services.UserService
	s3Service *clients.S3Service
}

type UserController interface {
	User(ctx *fiber.Ctx) error
	ChangePassword(ctx *fiber.Ctx) error
	AddProfile(ctx *fiber.Ctx) error
	UploadProfileImage(ctx *fiber.Ctx) error
	UpdateProfile(ctx *fiber.Ctx) error
}

func NewUserController(userService *services.UserService, s3Service *clients.S3Service) UserController {
	return &userController{userService: userService, s3Service: s3Service}
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

func (c userController) ChangePassword(ctx *fiber.Ctx) error {
	body := new(models.ChangePasswordRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	jwtClaims := ctx.Locals("user").(utils.UserCredential)
	model, err := c.userService.GetByUuid(jwtClaims.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Pengguna tidak ditemukan",
		})
	}

	model, err = c.userService.ChangePassword(model, body)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Gagal mengupdate password",
		})
	}

	response := models.FilterUserResponseV1(model)

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: response,
		Success: true,
		Message: "Sukses update password",
	})
}

func (c userController) UploadProfileImage(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Failed to parse form data",
		})
	}

	srcFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Failed to open uploaded file",
		})
	}
	defer func() {
		if srcFileErr := srcFile.Close(); srcFileErr != nil {
			log.Printf("Error closing file: %v", srcFileErr)
		}
	}()

	result, err := c.s3Service.UploadFileToS3(&srcFile, file)
	if err != nil {
		log.Printf("Error closing file: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Error uploading file",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data:    result,
		Success: true,
		Message: "Success Upload image",
	})
}

func (c userController) AddProfile(ctx *fiber.Ctx) error {
	body := new(models.UserProfileRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	jwtClaims := ctx.Locals("user").(utils.UserCredential)
	model, err := c.userService.GetByUuid(jwtClaims.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Pengguna tidak ditemukan",
		})
	}

	profile := models.WriteToModelUserProfile(body, model)

	result, err := c.userService.AddUserProfile(profile)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Gagal menambahkan profile",
		})
	}

	res := models.FilterUserProfileResponseV1(result, model)

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: res,
		Success: true,
		Message: "Sukses menambahkan profil pengguna",
	})
}

func (c userController) UpdateProfile(ctx *fiber.Ctx) error {
	body := new(models.UserProfileRequest)
	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Permintaan data tidak valid",
		})
	}

	jwtClaims := ctx.Locals("user").(utils.UserCredential)
	model, err := c.userService.GetByUuid(jwtClaims.ID)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Pengguna tidak ditemukan",
		})
	}

	profile := models.WriteToModelUserProfile(body, model)

	result, err := c.userService.UpdateUserProfile(profile)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(common.GlobalErrorHandlerResp{
			Success: false,
			Message: "Gagal mengubah profil",
		})
	}

	res := models.FilterUserProfileResponseV1(result, model)

	return ctx.Status(fiber.StatusOK).JSON(common.SuccessHandlerResp{
		Data: res,
		Success: true,
		Message: "Sukses mengubah profil pengguna",
	})
}