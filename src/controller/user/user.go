package userController

import (
	userRepository "pam/src/repository/user"

	"github.com/gin-gonic/gin"
)

type UserGateway interface {
	GetUsers(ctx *gin.Context)
}

type UserController struct {
	repository userRepository.UserGateway
}

func NewUserController(repository userRepository.UserGateway) UserGateway {
	return &UserController{repository: repository}
}
