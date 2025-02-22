package entity

import "time"

type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
}
