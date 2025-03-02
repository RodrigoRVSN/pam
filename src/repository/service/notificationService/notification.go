package notificationService

import (
	"time"
)

type NotificationGateway interface {
	CreateNotificationAsync(time time.Time, name string) error
}

type NotificationService struct{}

func NewNotificationService() NotificationGateway {
	return &NotificationService{}
}
