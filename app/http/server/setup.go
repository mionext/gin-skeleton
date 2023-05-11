package server

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	r "mionext/srv/app/http"
	"mionext/srv/app/http/middlewares"
	"net/http"
	"time"
)

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "The incorrect API route."})
}

func Setup() *http.Server {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()

	engine := gin.Default()
	engine.Use(middlewares.SrvTags())
	engine.NoRoute(notFound)
	engine.Use(requestid.New())
	engine.Use(gin.Recovery())
	r.Setup(engine)

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

	return &http.Server{
		Addr:           viper.GetString("http.listen"),
		Handler:        engine,
		ReadTimeout:    time.Second * 30,
		WriteTimeout:   time.Second * 30,
		MaxHeaderBytes: 1024 * 5,
	}
}
