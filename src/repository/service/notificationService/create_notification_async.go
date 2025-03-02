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

func createClient() (*eventbridge.Client, error) {
	cfg, cfgError := config.LoadDefaultConfig(context.TODO(),
		config.WithDefaultRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_ACCESS_SECRET"), "")))
	if cfgError != nil {
		return nil, cfgError
	}
	client := eventbridge.NewFromConfig(cfg)
	return client, nil
}

func createRule(client *eventbridge.Client, name string, time time.Time) error {
	cron := fmt.Sprintf("cron(%d %d %d %d ? %d)", time.Minute(), time.Hour(), time.Day(), time.Month(), time.Year())
	_, ruleError := client.PutRule(
		context.TODO(),
		&eventbridge.PutRuleInput{
			Name:               aws.String(name),
			ScheduleExpression: aws.String(cron),
		})
	return ruleError
}

func createTarget(client *eventbridge.Client, name string) error {
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

func (*NotificationService) CreateNotificationAsync(time time.Time, name string) error {
	client, cfgError := createClient()
	if cfgError != nil {
		return cfgError
	}
	ruleError := createRule(client, name, time)
	if ruleError != nil {
		return ruleError
	}
	targetError := createTarget(client, name)
	return targetError
}
