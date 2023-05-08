package middlewares

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	j "github.com/json-iterator/go"
)

var json = j.ConfigCompatibleWithStandardLibrary

type mioNextResponse struct {
	gin.ResponseWriter
	Body []byte
}

func (r *mioNextResponse) Write(b []byte) (int, error) {
	r.Body = append(r.Body, b...)
	return r.ResponseWriter.Write(b)
}

func Demo() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw := &mioNextResponse{
			ResponseWriter: c.Writer,
			Body:           []byte(""),
		}

		c.Writer = mw
		c.Next()

		rid := map[string]string{"rid": requestid.Get(c)}
		if b, err := json.Marshal(rid); err == nil {
			mw.Body = append(mw.Body, b...)
		}

		c.Writer.Write(mw.Body)
		mw.Body = nil
	}
}
