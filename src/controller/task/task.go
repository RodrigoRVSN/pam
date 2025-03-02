package taskController

import (
	"pam/src/repository/service/notificationService"
	taskRepository "pam/src/repository/task"

	"github.com/gin-gonic/gin"
)

type TaskGateway interface {
	GetTasks(ctx *gin.Context)
	CreateTask(ctx *gin.Context)
}

type TaskController struct {
	taskRepository      taskRepository.TaskGateway
	notificationService notificationService.NotificationGateway
}

func NewTaskController(taskRepository taskRepository.TaskGateway, notificationService notificationService.NotificationGateway) TaskGateway {
	return &TaskController{taskRepository: taskRepository, notificationService: notificationService}
}
