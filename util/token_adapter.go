package util

import (
	"fmt"
	"log"

	fibercasbinrest "github.com/prongbang/fiber-casbinrest"
	"github.com/riyan-eng/api-auth/middleware"
)

type tokenAdapter struct {
}

func NewTokenAdapter() fibercasbinrest.Adapter {
	return &tokenAdapter{}
}

func (r *tokenAdapter) GetRoleByToken(bearToken string) ([]string, error) {
	fmt.Println(bearToken)
	metaData, err := middleware.ExtractTokenMetadata("Bearer " + bearToken)
	if err != nil {
		log.Fatal(err.Error())
	}

	var roles []string
	for _, val := range metaData.Role {
		roles = append(roles, val.(string))
	}
	return roles, nil
}
