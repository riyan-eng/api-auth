package service

import (
	"errors"

	"github.com/riyan-eng/api-auth/middleware"
	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/riyan-eng/api-auth/module/management/repository"
	"github.com/riyan-eng/api-auth/module/management/service/entity"
	"github.com/riyan-eng/api-auth/util"
	"github.com/valyala/fasthttp"
)

type AuthService interface {
	Register(*dto.RegisterReq) error
	Login(*fasthttp.RequestCtx, *dto.LoginReq) (*entity.LoginEntity, error)
	Logout() error
}

type userController struct {
	User repository.UserInterface
}

func NewAuthService(repository repository.UserInterface) AuthService {
	return &userController{
		User: repository,
	}
}

func (repo *userController) Register(body *dto.RegisterReq) error {
	// entity
	entityReq := new(entity.Register)

	// generate hash password
	passwordHash := util.HashPassword(body.Password)
	entityReq.UserName = body.UserName
	entityReq.Password = passwordHash

	// communicate repo
	err := repo.User.Create(entityReq)
	if err != nil {
		return err
	}

	return nil
}

func (repo *userController) Login(ctx *fasthttp.RequestCtx, body *dto.LoginReq) (*entity.LoginEntity, error) {
	// entity
	entity := new(entity.LoginEntity)

	// communicate repository
	user, err := repo.User.GetUser(ctx, body)
	if err != nil {
		return nil, err
	}

	// verify password
	if !util.VerifyPassword(user.Password, body.Password) {
		return nil, errors.New("authentication failed")
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
