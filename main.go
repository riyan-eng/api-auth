package main

import (
	"log"
	"os"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
	"github.com/riyan-eng/api-auth/config"
	"github.com/riyan-eng/api-auth/module/management"
	"github.com/riyan-eng/api-auth/util"
)

func init() {
	os.Setenv("TZ", "Etc/UTC")
	config.LoadEnvironment()
	config.DatabaseConnection()
}

func main() {
	adapter := util.NewTokenAdapter()
	enforce, err := casbin.NewEnforcer("rbac_model.conf", "rbac_policy.csv")
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()
	app.Use(logger.New())

	management.Setup(app)

	// rbac middleware
	app.Use(fibercasbinrest.New(enforce, adapter))

	app.Listen(os.Getenv("SERVER_URL"))
}
