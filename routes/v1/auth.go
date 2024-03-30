package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/controllers"
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/middlewares"
	"github.com/rartstudio/gocourses/models"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/services"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SetupRoutesAuthV1(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config, db *gorm.DB, mail *gomail.Dialer, redis *redis.Client) {
	// service 
	userRepository := repositories.NewUserRepository(db)
	otpService := services.NewOtpService(config, redis, mail)
	jwtService := services.NewJWTService(config, redis)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(config, userRepository, otpService, jwtService, userService)
	
	// controller
	authController := controllers.NewAuthController(authService, userService)

	go otpService.HandleEmails()

	apiV1 := app.Group("/api/v1/auth")

	// ordinary route
	apiV1.Post("/register", func(c *fiber.Ctx) error {
		body := new(models.RegisterRequest)
		return common.ValidateRequest(c, customValidator, body)
	}, authController.Register)
	apiV1.Post("/login", func(c *fiber.Ctx) error {
		body := new(models.LoginRequest)
		return common.ValidateRequest(c, customValidator, body)
	}, authController.Login)

	// protected route
	apiV1.Use(middlewares.NewAuthMiddleware(config.JWTSECRET, jwtService))
	apiV1.Post("/verify", func(c *fiber.Ctx) error {
		body := new(models.VerifyAccountRequest)
		return common.ValidateRequest(c, customValidator, body)
	}, authController.Verify)
	apiV1.Post("/otp", authController.Otp)
}