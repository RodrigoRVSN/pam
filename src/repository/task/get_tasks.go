package taskRepository

import (
	"pam/src/domain/entity"
)

func (r *TaskRepository) GetTasks() ([]entity.Task, error) {
	rows, queryError := r.DB.Query("SELECT * FROM Tasks")
	if queryError != nil {
		panic(queryError.Error())
	}

	var tasks []entity.Task

	for rows.Next() {
		var task entity.Task
		if error := rows.Scan(&task.Id, &task.Title, &task.Description, &task.UserId, &task.DueDate); error != nil {
			return nil, error
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
