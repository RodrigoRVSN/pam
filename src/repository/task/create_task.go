package taskRepository

import (
	"context"
	"pam/src/domain/entity"
)

func (r *TaskRepository) CreateTask(task entity.Task) (int64, error) {
	result, queryError := r.DB.ExecContext(context.Background(), "INSERT INTO Tasks (title, description, due_date, user_id) VALUES (?, ?, ?, ?)", task.Title, task.Description, task.DueDate, task.UserId)
	if queryError != nil {
		return 0, queryError
	}
	lastId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return lastId, nil
}
