package service

import (
	"fmt"

	"github.com/riyan-eng/api-auth/middleware"
	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/riyan-eng/api-auth/module/management/repository"
	"github.com/valyala/fasthttp"
)

type AuthService interface {
	Login(*fasthttp.RequestCtx, *dto.LoginReq) error
	Logout() error
}

type userController struct {
	repo repository.UserInterface
}

func NewAuthService(repository repository.UserInterface) AuthService {
	return &userController{
		repo: repository,
	}
}

func (repo *userController) Login(ctx *fasthttp.RequestCtx, body *dto.LoginReq) error {
	fmt.Println(body)
	err := repo.repo.GetUser(ctx, body)

	token, err := middleware.GenerateNewAccessToken()
	fmt.Println(token)
	return err
}

func (repo *userController) Logout() error {
	return nil
}
