package service

import (
	"github.com/riyan-eng/api-auth/module/management/repository"
	"github.com/valyala/fasthttp"
)

type AuthService interface {
	Login(*fasthttp.RequestCtx) error
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

func (repo *userController) Login(ctx *fasthttp.RequestCtx) error {
	err := repo.repo.GetUser(ctx)
	return err
}

func (repo *userController) Logout() error {
	return nil
}
