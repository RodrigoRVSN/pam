package main

import (
	taskController "pam/src/controller/task"
	userController "pam/src/controller/user"
	"pam/src/infra/db"
	taskRepository "pam/src/repository/task"
	userRepository "pam/src/repository/user"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	error := godotenv.Load()
	if error != nil {
		panic(error.Error())
	}
	engine := gin.Default()
	db := db.InitializeDatabase()

	userRepository := userRepository.NewUserRepository(db)
	userController := userController.NewUserController(userRepository)

	taskRepository := taskRepository.NewTaskRepository(db)
	taskController := taskController.NewTaskController(taskRepository)

	engine.GET("/users", userController.GetUsers)
	engine.POST("/create-user", userController.CreateUser)

	engine.GET("/tasks", taskController.GetTasks)
	engine.POST("/create-task", taskController.CreateTask)

	engine.Run()
}
