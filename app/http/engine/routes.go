package engine

import (
	"github.com/gin-gonic/gin"
	authHandler "mionext/srv/app/http/handlers/auth"
)

func mountRoutes(engine *gin.Engine) *gin.Engine {
	auth := engine.Group("auth")
	{
		auth.GET("sign-in", authHandler.SignIn)
	}

	return engine
}
