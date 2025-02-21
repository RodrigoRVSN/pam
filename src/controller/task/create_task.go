package taskController

import (
	"context"
	"fmt"
	"net/http"
	"pam/src/domain/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
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

	cfg, error := config.LoadDefaultConfig(context.TODO(), config.WithDefaultRegion("us-east-1"))
	client := eventbridge.NewFromConfig(cfg)
	output, error := client.PutRule(context.TODO(), &eventbridge.PutRuleInput{Name: aws.String(fmt.Sprintf("%d", newId))})
	fmt.Println(output)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
	}
	ctx.JSON(http.StatusCreated, newId)
}
