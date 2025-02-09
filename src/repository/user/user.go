package userRepository

import (
	"database/sql"
	"pam/src/domain/entity"
)

type UserRepository struct {
	DB *sql.DB
}

type UserGateway interface {
	GetUsers() ([]entity.User, error)
	CreateUser(user entity.User) (int64, error)
}

func NewUserRepository(db *sql.DB) UserGateway {
	return &UserRepository{DB: db}
}
