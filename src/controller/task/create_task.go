package taskController

import (
	"fmt"
	"net/http"
	"pam/src/domain/entity"
	"time"

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

	reminderTime := task.DueDate.Add(-time.Duration(15) * time.Minute)
	go c.notificationService.CreateNotificationAsync(reminderTime, fmt.Sprintf("%d", newId))
	ctx.JSON(http.StatusCreated, newId)
}
