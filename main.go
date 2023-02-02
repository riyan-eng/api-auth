package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-auth/config"
)

func init() {
	config.LoadEnvironment()
	config.DatabaseConnection()
}

func main() {
	app := fiber.New()

	app.Listen(os.Getenv("SERVER_URL"))
}
