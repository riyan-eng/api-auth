package main

import (
	"fmt"
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

	var number int = 1

	fmt.Println(number)

	app.Use(logger.New())

	app.Listen(os.Getenv("SERVER_URL"))
}
