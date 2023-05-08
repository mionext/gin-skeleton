package health

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK", "rid": requestid.Get(c)})
}

func Stats(c *gin.Context) {
	stats := gin.H{
		"message": "OK", "rid": requestid.Get(c),
	}

	c.JSON(http.StatusOK, stats)
}
