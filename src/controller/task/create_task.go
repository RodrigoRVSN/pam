package taskController

import (
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
	tasks, error := c.taskRepository.CreateTask(task)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
