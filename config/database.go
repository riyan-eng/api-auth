package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/riyan-eng/api-auth/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	// declare variable
	var err error
	var connectionPool *sql.DB

	// define db connection setting
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// database dns
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// database conn
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		fmt.Printf("error, %s", err)
		panic("not connected to database")
	}

	// connection pool
	connectionPool, err = DB.DB()
	connectionPool.SetMaxOpenConns(maxConn)
	connectionPool.SetMaxIdleConns(maxIdleConn)
	connectionPool.SetConnMaxLifetime(time.Duration(maxLifetimeConn))
	if err != nil {
		fmt.Println("error, create connection pool")
	}

	if err := connectionPool.Ping(); err != nil {
		fmt.Println("error, not sent ping to database")
	}

	// migration
	DB.AutoMigrate(&util.User{})

	fmt.Println("connected to database", os.Getenv("DB_NAME"))
}
