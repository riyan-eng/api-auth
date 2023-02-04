package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/riyan-eng/api-auth/middleware"
	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/riyan-eng/api-auth/module/management/service"
	"github.com/riyan-eng/api-auth/util"
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
	authRoute.Post("/logout", middleware.JWTProtected(), s.Logout)
}

func (service authService) Login(c *fiber.Ctx) error {
	// parse body
	var body dto.LoginReq
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}

	// validate body
	if err := util.Validate(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"data":    err,
			"message": "bad",
		})
	}

	// communicate service
	if err := service.service.Login(c.Context(), &body); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"data":    err.Error(),
			"message": "bad",
		})
	}

	// response ok
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "login",
		"message": "ok",
	})
}

func (service authService) Logout(c *fiber.Ctx) error {
	bearToken := c.Get("Authorization")
	data, err := middleware.ExtractTokenMetadata(bearToken)

	fmt.Println(data)
	fmt.Println(err)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    "logout",
		"message": "ok",
	})
}
