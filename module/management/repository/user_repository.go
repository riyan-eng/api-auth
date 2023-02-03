package repository

import (
	"database/sql"
	"fmt"

	"github.com/valyala/fasthttp"
)

type UserInterface interface {
	GetUser(*fasthttp.RequestCtx) error
}

type database struct {
	Db *sql.DB
}

func NewUserInterface(DB *sql.DB) UserInterface {
	return &database{
		Db: DB,
	}
}

func (db *database) GetUser(ctx *fasthttp.RequestCtx) error {

	var name string
	fmt.Println("jjj")
	rows := db.Db.QueryRowContext(ctx, "select name from management.users where name='RIYAN'").Scan(&name)
	// if err != nil {
	// 	return err
	// }

	fmt.Println(name)
	fmt.Println(rows)
	return nil
}
