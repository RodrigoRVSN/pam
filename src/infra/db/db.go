package db

import (
	"database/sql"
	"os"
	"time"
)

func InitializeDatabase() *sql.DB {
	db, error := sql.Open("mysql", os.Getenv("DB_DSN_NAME"))
	if error != nil {
		panic(error)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
