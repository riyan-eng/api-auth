package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type UserInterface interface {
	GetUser() error
}

type database struct {
	Db *sql.DB
}

func NewUserInterface(DB *sql.DB) UserInterface {
	return &database{
		Db: DB,
	}
}

func (db *database) GetUser() error {

	type User struct {
		Name string
	}
	fmt.Println("jjj")
	rows := db.Db.QueryRowContext(context.Background(), "select * from management.users where name=RIYAN").Scan(&User{})
	// if err != nil {
	// 	return err
	// }

	fmt.Println(rows)
	return nil
}
