package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/riyan-eng/api-auth/module/management/controller/dto"
	"github.com/riyan-eng/api-auth/module/management/repository/model"
	"github.com/riyan-eng/api-auth/module/management/service/entity"
	"github.com/valyala/fasthttp"
)

type UserInterface interface {
	GetUser(*fasthttp.RequestCtx, *dto.LoginReq) (*model.User, error)
	Create(*entity.Register) error
}

type database struct {
	Db *sql.DB
}

func NewUserInterface(DB *sql.DB) UserInterface {
	return &database{
		Db: DB,
	}
}

func (db *database) Create(entityReq *entity.Register) error {
	query := fmt.Sprintf(`
		insert into management.users(name, password) values('%v', '%v')
	`, entityReq.UserName, entityReq.Password)

	_, err := db.Db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (db *database) GetUser(ctx *fasthttp.RequestCtx, body *dto.LoginReq) (*model.User, error) {
	user := new(model.User)
	query := fmt.Sprintf(`
		select id, name, password from management.users where name='%v'
	`, body.UserName)

	err := db.Db.QueryRowContext(ctx, query).Scan(&user.ID, &user.Name, &user.Password)
	if err == sql.ErrNoRows {
		return nil, errors.New("no data")
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
