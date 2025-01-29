package main

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int64  `json:"id"`
}

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Id          int64  `json:"id"`
	UserId      int64  `json:"user_id"`
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
				c.JSON(http.StatusBadRequest, error.Error())
			}
			users = append(users, user)
		}
		c.JSON(http.StatusOK, users)
	})

	engine.GET("/tasks", func(c *gin.Context) {
		rows, queryError := db.Query("SELECT * FROM Tasks")
		if queryError != nil {
			panic(queryError.Error())
		}

		var tasks []Task

		for rows.Next() {
			var task Task
			if error := rows.Scan(&task.Id, &task.Title, &task.Description, &task.UserId, &task.DueDate); error != nil {
				c.JSON(http.StatusBadRequest, error.Error())
			}
			tasks = append(tasks, task)
		}
		c.JSON(http.StatusOK, tasks)
	})

	engine.POST("/create-task", func(c *gin.Context) {
		var task Task
		if error := c.ShouldBindJSON(&task); error != nil {
			c.JSON(http.StatusBadRequest, error.Error())
		}
		result, queryError := db.ExecContext(context.Background(), "INSERT INTO Tasks (title, description, due_date, user_id) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.DueDate, task.UserId)
		if queryError != nil {
			panic(queryError.Error())
		}
		lastId, error := result.LastInsertId()
		if error != nil {
			panic(error.Error())
		}
		c.JSON(http.StatusOK, lastId)
	})

	engine.Run()
}
