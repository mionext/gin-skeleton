package middlewares

import (
	"github.com/gin-gonic/gin"
	"os"
)

func SrvTags() gin.HandlerFunc {
	name, _ := os.Hostname()
	return func(c *gin.Context) {
		c.Header("X-Who", name)

		c.Next()
	}
}
