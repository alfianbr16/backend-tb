package main

import (
	"log"
	"os"

	"github.com/alfianbr16/backend-tb/config"
	"github.com/alfianbr16/backend-tb/url"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "github.com/alfianbr16/backend-tb/docs"

)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/alfianbr16
// @contact.email 714220048@std.ulbi.ac.id

// @host backend-tb-97f8facdbf4c.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	app := fiber.New()

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New(config.Cors))

	// Routes
	url.Web(app)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).SendString("Sorry, page not found!")
	})

	// Get port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		port = "3200"
	}

	// Start server
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}
