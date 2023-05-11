package http

import (
	"github.com/gin-gonic/gin"
	authHandler "mionext/srv/app/http/handlers/auth"
)

func Setup(engine *gin.Engine) *gin.Engine {
	auth := engine.Group("auth")
	{
		auth.GET("sign-in", authHandler.SignIn)
	}

	return engine
}
