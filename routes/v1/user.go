package v1

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/clients"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/controllers"
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/middlewares"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/services"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func SetupRoutesUserV1(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config, db *gorm.DB, mail *gomail.Dialer, redis *redis.Client) {
	// connect to s3
	s3Client, err := initializers.ConnectToS3(config)
	if err != nil {
		log.Fatalf("Failed to initialize S3 client: %v", err)
	}

	// service
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	s3Service := clients.NewS3Service(s3Client, config)
	jwtService := services.NewJWTService(config, redis)
	
	// controller 
	userController := controllers.NewUserController(userService, s3Service)

	apiV1 := app.Group("/api/v1/users")

	// protected route
	apiV1.Use(middlewares.NewAuthMiddleware(config.JWTSECRET, jwtService))

	apiV1.Get("/", userController.GetUser)
	apiV1.Put("/change-password", userController.ChangePassword)
	apiV1.Post("/upload-profile-image", userController.UploadProfileImage)
	apiV1.Post("/profile", userController.AddProfile)
	apiV1.Put("/profile", userController.UpdateProfile)
} 