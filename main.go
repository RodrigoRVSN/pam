package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	CreatedAt string `json:"created_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int64  `json:"id"`
}

// TODO: reimplement the python api in golang and remove python implementation

func main() {
	engine := gin.Default()
	db, error := sql.Open("mysql", "root:password@/task_management")
	if error != nil {
		panic(error)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	engine.GET("/users", func(c *gin.Context) {
		rows, queryError := db.Query("SELECT * FROM Users")
		if queryError != nil {
			panic(queryError.Error())
		}

		var users []User

		for rows.Next() {
			var user User
			if error := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.Password); error != nil {
				fmt.Println(error.Error())
			}
			users = append(users, user)
		}
		fmt.Println(users)
		c.JSON(http.StatusOK, users)
	})

	engine.Run()
}
