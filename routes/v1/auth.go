package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/controllers"
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/services"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SetupRoutesAuthV1(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config, db *gorm.DB, mail *gomail.Dialer, redis *redis.Client) {
	// service 
	authRepository := repositories.NewUserRepository(db)
	otpService := services.NewOtpService(config, redis, mail)
	jwtService := services.NewJWTService(config, redis)
	authService := services.NewAuthService(config, authRepository, otpService, jwtService)
	authController := controllers.NewAuthController(authService)

	apiV1 := app.Group("/api/v1/auth")
	apiV1.Post("/register", func(c *fiber.Ctx) error {
		body := new(models.RegisterRequest)
		return common.ValidateRequest(c, customValidator, body)
	}, authController.Register)
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