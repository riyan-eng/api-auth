package config

import (
	"fmt"

	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
	"github.com/riyan-eng/api-auth/middleware"
)

type tokenAdapter struct {
}

func NewTokenAdapter() fibercasbinrest.Adapter {
	return &tokenAdapter{}
}

func (r *tokenAdapter) GetRoleByToken(reqToken string) ([]string, error) {
	fmt.Println(reqToken)
	metaData, _ := middleware.ExtractTokenMetadata(reqToken)
	fmt.Println(metaData)
	role := "admin"
	return []string{role}, nil
}
