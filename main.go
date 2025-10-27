package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"foliage/config"
	"foliage/routes"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	config.ConnectDB()
	routes.SetupRoutes(app)
	app.Static("/", "./public")
	app.Static("/uploads", "./uploads")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen(":" + port))
}
