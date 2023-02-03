package service

import "github.com/riyan-eng/api-auth/module/management/repository"

type AuthService interface {
	Login() error
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

func (repo *userController) Login() error {
	err := repo.repo.GetUser()
	return err
}

func (repo *userController) Logout() error {
	return nil
}
