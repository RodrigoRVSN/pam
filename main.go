package main

import (
	"context"
	"net/http"
	taskController "pam/src/controller/task"
	userController "pam/src/controller/user"
	"pam/src/domain/entity"
	"pam/src/infra/db"
	taskRepository "pam/src/repository/task"
	userRepository "pam/src/repository/user"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	engine := gin.Default()
	db := db.InitializeDatabase()
	userRepository := userRepository.NewUserRepository(db)
	userController := userController.NewUserController(userRepository)
	taskRepository := taskRepository.NewTaskRepository(db)
	taskController := taskController.NewTaskController(taskRepository)

	engine.GET("/users", userController.GetUsers)
	engine.GET("/tasks", taskController.GetTasks)
	engine.POST("/create-task", taskController.CreateTask)

	engine.POST("/create-user", func(c *gin.Context) {
		var user entity.User
		if error := c.ShouldBindJSON(&user); error != nil {
			c.JSON(http.StatusBadRequest, error.Error())
		}
		query := "INSERT INTO Users (name, email, password, created_at) VALUES (?, ?, ?, ?)"
		result, queryError := db.ExecContext(context.Background(), query, user.Name, user.Email, user.Password, time.Now())
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
