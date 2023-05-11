package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	errors "mionext/srv/app/errors"
	"net/http"
)

var (
	OK = gin.H{"message": "OK"}
)

func ValidationError(c *gin.Context, err error) {
	errs := []string{}
	if err, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range err {
			err := &errors.FieldErrorWrapper{fieldError}
			errStr := err.Error()
			if p := err.Param(); p != "" {
				errStr += fmt.Sprintf(" = %s", p)
			}

			errs = append(errs, errStr)
		}
	} else {
		errs = append(errs, "Validation error.")
	}

	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": errs})
}

func InternalError(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Internal server error."})
}
