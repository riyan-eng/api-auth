package management

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-auth/config"
	"github.com/riyan-eng/api-auth/module/management/controller"
	"github.com/riyan-eng/api-auth/module/management/repository"
	"github.com/riyan-eng/api-auth/module/management/service"
)

func Setup(app *fiber.App) {
	repo := repository.NewUserInterface(config.DB)
	svc := service.NewAuthService(repo)
	controller.NewAuthController(svc, app)
}
