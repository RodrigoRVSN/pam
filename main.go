package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	created_at string
	name       string
	email      string
	password   string
	id         int64
}

// TODO: reimplement the python api in golang and remove python implementation

func main() {
	db, error := sql.Open("mysql", "root:password@/task_management")
	if error != nil {
		panic(error)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, queryError := db.Query("SELECT * FROM Users")
	if queryError != nil {
		panic(queryError.Error())
	}

	var users []User

	for rows.Next() {
		var user User
		if error := rows.Scan(&user.id, &user.name, &user.email, &user.created_at, &user.password); error != nil {
			fmt.Println(error.Error())
		}
		users = append(users, user)
	}
	fmt.Println(users)
}
