package taskController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *TaskController) GetTasks(ctx *gin.Context) {
	tasks, error := c.taskRepository.GetTasks()
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}
