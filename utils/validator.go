package utils

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func FormatValidationError(err error) []ResponseError {
	var result []ResponseError

	if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
		for _, fe := range ve {
			result = append(result, ResponseError{
				Field:   fe.Field(),
				Message: msgForTag(fe),
			})
		}
	}
	return result
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "min":
		return fmt.Sprintf("Minimum length is %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("Maximum length is %s characters", fe.Param())
	default:
		return "Invalid value"
	}
}
