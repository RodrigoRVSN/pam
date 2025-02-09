package taskRepository

import (
	"database/sql"
	"pam/src/domain/entity"
)

type TaskRepository struct {
	DB *sql.DB
}

type TaskGateway interface {
	GetTasks() ([]entity.Task, error)
}

func NewTaskRepository(db *sql.DB) TaskGateway {
	return &TaskRepository{DB: db}
}
