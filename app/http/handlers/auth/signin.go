package auth

import (
	"github.com/gin-gonic/gin"
	"mionext/srv/app/service/jwt"
	"net/http"
)

func SignIn(c *gin.Context) {
	token, err := jwt.New().IssueById(1)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
