package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DatabaseConnection() {
	// define db connection setting
	maxConn, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION"))
	maxIdleConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	maxLifetimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))

	// database dns
	dSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// database conn
	DB, _ = sql.Open("postgres", dSN)

	// connection pool
	DB.SetMaxOpenConns(maxConn)
	DB.SetMaxIdleConns(maxIdleConn)
	DB.SetConnMaxLifetime(time.Duration(maxLifetimeConn))

	if err := DB.Ping(); err != nil {
		log.Fatal("error, not connected to database")
	}

	fmt.Println("connected to database", os.Getenv("DB_NAME"))
}
