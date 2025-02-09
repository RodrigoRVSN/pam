package taskController

import (
	taskRepository "pam/src/repository/task"

	"github.com/gin-gonic/gin"
)

type TaskGateway interface {
	GetTasks(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
}

type TaskController struct {
	taskRepository taskRepository.TaskGateway
}

func NewTaskController(taskRepository taskRepository.TaskGateway) TaskGateway {
	return &TaskController{taskRepository: taskRepository}
}
