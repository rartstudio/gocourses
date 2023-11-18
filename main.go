package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rartstudio/gocourses/initializers"
	"github.com/rartstudio/gocourses/routes"
)

func main() {
	// load environment
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load environment variables \n", err.Error())
	}

	// connect to database
	db := initializers.ConnectDB(&config);

	// init app with config
	app := fiber.New(fiber.Config{
		AppName: config.APPNAME,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// using logger
	app.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} - ${method} ${path}\n"}))

	// health check route
	app.Get("/health-checker", func(c *fiber.Ctx) error {
		dbStatus := true
		dbMessage := "Success connect to Database"

		err := db.Exec("SELECT 1").Error

		if err != nil {
			dbStatus = false
			dbMessage = "Error connect to Database"
		}
		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"app": fiber.Map{
				"status": true,
				"message": "Application is up",
			},
			"database": fiber.Map{
				"status": dbStatus,
				"message": dbMessage,
			},
		})
	})

	routes.SetupRoutes(app)
	
	port := fmt.Sprintf(":%d", config.PORT)
	err = app.Listen(port)
	if err != nil {
		log.Fatalf("Error running app: %v", err)
	}

	fmt.Printf("App running on port %s \\n", port)
	
}