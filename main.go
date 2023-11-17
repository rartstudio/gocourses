package main

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rartstudio/gocourses/initializers"
)

func main() {
// load environment
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load environment variables \n", err.Error())
	}

	// init app with config
	app := fiber.New(fiber.Config{
		AppName: config.APPNAME,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	// using logger
	app.Use(logger.New(logger.Config{Format: "[${ip}]:${port} ${status} - ${method} ${path}\n"}))

	// health check route
	
	port := fmt.Sprintf(":%d", config.PORT)
	err = app.Listen(port)
	if err != nil {
		log.Fatalf("Error running app: %v", err)
	}

	fmt.Printf("App running on port %s \\n", port)
	
}