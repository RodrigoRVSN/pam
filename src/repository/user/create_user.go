package userRepository

import (
	"context"
	"pam/src/domain/entity"
	"time"
)

func (r *UserRepository) CreateUser(user entity.User) (int64, error) {
	query := "INSERT INTO Users (name, email, password, created_at) VALUES (?, ?, ?, ?)"
	result, queryError := r.DB.ExecContext(context.Background(), query, user.Name, user.Email, user.Password, time.Now())
	if queryError != nil {
		return 0, queryError
	}
	lastId, error := result.LastInsertId()
	if error != nil {
		return 0, error
	}
	return lastId, nil
}
