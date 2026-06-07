package httpinfra

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var instance = validator.New()

func Validate(s any) error {
	return instance.Struct(s)
}

func ParseErrors(err error) map[string]string {
	result := make(map[string]string)

	if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
		for _, fe := range ve {
			result[fe.Field()] = fe.Tag()
		}
	}

	return result
}
