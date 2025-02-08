package userRepository

import "database/sql"

type User struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int64  `json:"id"`
}

type UserRepository struct {
	DB *sql.DB
}

type UserGateway interface {
	GetUsers() ([]User, error)
}

func NewUserRepository(db *sql.DB) UserGateway {
	return &UserRepository{DB: db}
}
