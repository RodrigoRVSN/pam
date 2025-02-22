package taskController

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"pam/src/domain/entity"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
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

	cfg, error := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), "")))
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
	}
	client := eventbridge.NewFromConfig(cfg)
	cron := fmt.Sprintf("cron(%d %d %d %d ? %d)", task.DueDate.Minute(), task.DueDate.Hour(), task.DueDate.Day(), task.DueDate.Month(), task.DueDate.Day())

	_, err := client.PutRule(
		context.TODO(),
		&eventbridge.PutRuleInput{
			Name:               aws.String(fmt.Sprintf("%d", newId)),
			ScheduleExpression: aws.String(cron),
		})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
	}
	ctx.JSON(http.StatusCreated, newId)
}
