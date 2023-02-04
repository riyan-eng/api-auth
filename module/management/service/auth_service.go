package service

import (
	"github.com/riyan-eng/api-auth/middleware"
	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/riyan-eng/api-auth/module/management/repository"
	"github.com/riyan-eng/api-auth/module/management/service/entity"
	"github.com/valyala/fasthttp"
)

type AuthService interface {
	Login(*fasthttp.RequestCtx, *dto.LoginReq) (*entity.LoginEntity, error)
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

func (repo *userController) Login(ctx *fasthttp.RequestCtx, body *dto.LoginReq) (*entity.LoginEntity, error) {
	// entity
	entity := new(entity.LoginEntity)

	// communicate repository
	user, err := repo.repo.GetUser(ctx, body)
	if err != nil {
		return nil, err
	}
	entity.Name = user.Name

	// communicate jwt middleware
	token, err := middleware.GenerateNewAccessToken()
	if err != nil {
		return nil, err
	}
	entity.AccessToken = token.Token
	entity.AccessTokenExpired = token.Expired

	// response
	return entity, nil
}

func (repo *userController) Logout() error {
	return nil
}
