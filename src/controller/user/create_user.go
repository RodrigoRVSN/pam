package userController

import (
	"net/http"
	"pam/src/domain/entity"

	"github.com/gin-gonic/gin"
)

func (c UserController) CreateUser(ctx *gin.Context) {
	var user entity.User
	if error := ctx.ShouldBindJSON(&user); error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}
	newId, error := c.repository.CreateUser(user)
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}
	ctx.JSON(http.StatusCreated, newId)
}
