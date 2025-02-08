package userController

import (
	userRepository "pam/repository"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Id        int64  `json:"id"`
}

type UserGateway interface {
	GetUsers(ctx *gin.Context)
}

type UserController struct {
	repository userRepository.UserGateway
}

func NewUserController(repository userRepository.UserGateway) UserGateway {
	return &UserController{repository: repository}
}
