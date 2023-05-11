package errors

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

const (
	fieldErrMsg = "Filed: %s validation failed for: %s"
)

type FieldErrorWrapper struct {
	validator.FieldError
}

func (fe *FieldErrorWrapper) Error() string {
	return fmt.Sprintf(fieldErrMsg, fe.Field(), fe.Tag())
}
