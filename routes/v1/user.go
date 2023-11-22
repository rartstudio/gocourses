package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rartstudio/gocourses/common"
	"github.com/rartstudio/gocourses/controllers"
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/middlewares"
	"github.com/rartstudio/gocourses/repositories"
	"github.com/rartstudio/gocourses/services"
	"gorm.io/gorm"
)

func SetupRoutesUserV1(app *fiber.App, customValidator *common.CustomValidator, config *initializers.Config, db *gorm.DB) {
	// service
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	
	// controller 
	userController := controllers.NewUserController(userService)

	apiV1 := app.Group("/api/v1/users")

	// protected route
	apiV1.Use(middlewares.NewAuthMiddleware(config.JWTSECRET))

	apiV1.Get("/user", userController.User)
} 