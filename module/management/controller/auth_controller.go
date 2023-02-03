package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-auth/module/management/service"
)

// type AuthController interface {
// 	Login(*fiber.Ctx) error
// 	Logout(*fiber.Ctx) error
// }

type authService struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService, route *fiber.App) {
	s := &authService{
		service: service,
	}
	authRoute := route.Group("/auth")
	authRoute.Post("/login", s.Login)
	authRoute.Post("/logout", s.Logout)
}

func (service authService) Login(c *fiber.Ctx) error {
	if err := service.service.Login(); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "login",
		"message": "ok",
	})
}

func (service authService) Logout(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "logout",
		"message": "ok",
	})
}
