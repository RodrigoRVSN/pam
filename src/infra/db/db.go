package db

import (
	"database/sql"
	"time"
)

func InitializeDatabase() *sql.DB {
	db, error := sql.Open("mysql", "root:password@tcp(db:3306)/task_management")
	if error != nil {
		panic(error)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
