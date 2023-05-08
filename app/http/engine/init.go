package engine

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"mionext/srv/app/http/handlers/health"
)

func Init() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	engine := gin.Default()
	engine.Use(requestid.New())
	engine.Use(gin.Recovery())
	if viper.GetBool("http.enable_cors") {
		//engine.Use(cors.New(cors.Config{
		//	AllowAllOrigins:        false,
		//	AllowOrigins:           nil,
		//	AllowOriginFunc:        nil,
		//	AllowMethods:           nil,
		//	AllowHeaders:           nil,
		//	AllowCredentials:       false,
		//	ExposeHeaders:          nil,
		//	MaxAge:                 0,
		//	AllowWildcard:          false,
		//	AllowBrowserExtensions: false,
		//	AllowWebSockets:        false,
		//	AllowFiles:             false,
		//}))
	}

	engine.GET("health", health.Check)

	return mountRoutes(engine)
}
