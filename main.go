package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/riyan-eng/api-auth/config"
)

func init() {
	config.LoadEnvironment()
	config.DatabaseConnection()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Listen(os.Getenv("SERVER_URL"))
}
