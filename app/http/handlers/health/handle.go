package health

import (
	"github.com/gin-gonic/gin"
	"mionext/srv/app/http/handlers"
	"net/http"
)

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, handlers.OK)
}

func Stats(c *gin.Context) {
	c.JSON(http.StatusOK, handlers.OK)
}
