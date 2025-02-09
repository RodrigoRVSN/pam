package userController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, error := c.repository.GetUsers()
	if error != nil {
		ctx.JSON(http.StatusBadRequest, error.Error())
		return
	}
	ctx.JSON(http.StatusOK, users)
}
