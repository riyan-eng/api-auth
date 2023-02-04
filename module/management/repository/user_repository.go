package repository

import (
	"database/sql"
	"fmt"

	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/valyala/fasthttp"
)

type UserInterface interface {
	GetUser(*fasthttp.RequestCtx, *dto.LoginReq) error
}

type database struct {
	Db *sql.DB
}

func NewUserInterface(DB *sql.DB) UserInterface {
	return &database{
		Db: DB,
	}
}

func (db *database) GetUser(ctx *fasthttp.RequestCtx, body *dto.LoginReq) error {
	var name string
	query := fmt.Sprintf(`
		select name from management.users where name='%v'
	`, body.UserName)
	// err := db.Db.QueryRowContext(ctx, query).Scan(&name)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(name)
	fmt.Println(query)
	// fmt.Println(rows)
	return nil
}
