package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// urlExample := "postgres://username:password@localhost:5432/database_name?sslmode=disable"
func Init(url string) {
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
