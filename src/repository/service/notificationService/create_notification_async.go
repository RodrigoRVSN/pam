package notificationService

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
)

func (*NotificationService) CreateNotificationAsync(time time.Time, name string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), "")))
	if err != nil {
		return err
	}
	client := eventbridge.NewFromConfig(cfg)
	cron := fmt.Sprintf("cron(%d %d %d %d ? %d)", time.Minute(), time.Hour(), time.Day(), time.Month(), time.Year())

	_, ruleError := client.PutRule(
		context.TODO(),
		&eventbridge.PutRuleInput{
			Name:               aws.String(name),
			ScheduleExpression: aws.String(cron),
		})
	if ruleError != nil {
		return ruleError
	}

	_, targetError := client.PutTargets(context.TODO(), &eventbridge.PutTargetsInput{
		Rule: aws.String(name),
		Targets: []types.Target{
			{
				Id:      aws.String("1"),
				Arn:     aws.String(os.Getenv("AWS_ARN_SNS_TASK_REMINDER")),
				RoleArn: aws.String(os.Getenv("AWS_ARN_SNS_TASK_REMINDER_ROLE")),
			},
		},
	})

	return targetError
}
