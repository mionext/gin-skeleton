package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mionext/srv/app/http/handlers"
	"mionext/srv/app/service/jwt"
	"net/http"
)

type (
	loginRequest struct {
		Mobile string `form:"mobile" json:"mobile" binding:"required,len=11"`
		Secret string `form:"password" json:"mobile" binding:"required,min=6"`
	}
)

func SignIn(c *gin.Context) {
	var lr loginRequest
	if err := c.ShouldBindWith(&lr, binding.Query); err != nil {
		handlers.ValidationError(c, err)
		return
	}

	token, err := jwt.New().IssueById(1)
	if err != nil {
		handlers.InternalError(c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
