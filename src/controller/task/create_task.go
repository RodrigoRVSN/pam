package taskController

import (
	"fmt"
	"net/http"
	"pam/src/domain/entity"

	"github.com/gin-gonic/gin"
)

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task entity.Task
	if error := ctx.ShouldBindJSON(&task); error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}

	newId, error := c.taskRepository.CreateTask(task)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}

	go c.notificationService.CreateNotificationAsync(task.DueDate, fmt.Sprintf("%d", newId))
	ctx.JSON(http.StatusCreated, newId)
}
