package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
)

type ResponseError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type FormattedError struct {
	Message string
	Errors  any
}

func FormatValidationError(err error) FormattedError {
	if errors.Is(err, io.EOF) || err.Error() == "EOF" {
		return FormattedError{
			Message: "Request body cannot be empty",
			Errors:  nil,
		}
	}

	if syntaxError, ok := errors.AsType[*json.SyntaxError](err); ok {
		return FormattedError{
			Message: "Invalid JSON format: " + syntaxError.Error(),
			Errors:  nil,
		}
	}

	if unmarshalTypeError, ok := errors.AsType[*json.UnmarshalTypeError](err); ok {
		return FormattedError{
			Message: "Invalid type for field: " + unmarshalTypeError.Field,
			Errors:  nil,
		}
	}

	if bindingError, ok := errors.AsType[*gin.Error](err); ok {
		err = bindingError.Err
	}

	if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
		var result []ResponseError
		for _, fe := range ve {
			result = append(result, ResponseError{
				Field:   fe.Field(),
				Message: msgForTag(fe),
			})
		}
		return FormattedError{
			Message: "Invalid validation",
			Errors:  result,
		}
	}

	if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {

		// "23505" is the global PostgreSQL error code for Unique Violation
		if pgErr.Code == "23505" {
			var result []ResponseError

			fieldName := "Database"
			message := "Duplicate entry detected on a unique field"

			if pgErr.ConstraintName != "" {
				parts := strings.Split(pgErr.ConstraintName, "_")
				rawField := parts[len(parts)-1]

				fieldName = rawField[:1] + rawField[1:]
				message = fmt.Sprintf("This %s is already taken", rawField)
			}

			result = append(result, ResponseError{
				Field:   fieldName,
				Message: message,
			})

			return FormattedError{
				Message: "Invalid validation",
				Errors:  result,
			}
		}
	}

	return FormattedError{
		Message: "An unexpected error occurred",
		Errors:  nil,
	}
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email address"
	case "url":
		return "Invalid URL format"
	case "alphanum":
		return "This field must only contain letters and numbers"
	case "alpha":
		return "This field must only contain letters"
	case "numeric":
		return "This field must only contain numbers"
	case "min":
		if fe.Kind() == reflect.String {
			return fmt.Sprintf("Minimum length is %s characters", fe.Param())
		}
		return fmt.Sprintf("Minimum value must be %s", fe.Param())
	case "max":
		if fe.Kind() == reflect.String {
			return fmt.Sprintf("Maximum length is %s characters", fe.Param())
		}
		return fmt.Sprintf("Maximum value must be %s", fe.Param())
	case "len":
		return fmt.Sprintf("This field must be exactly %s characters or items long", fe.Param())
	case "eqfield":
		return fmt.Sprintf("This field must match the %s field", fe.Param())
	case "nefield":
		return fmt.Sprintf("This field cannot be the same as the %s field", fe.Param())
	case "oneof":
		return fmt.Sprintf("This field must be one of the following values: %s", fe.Param())
	case "boolean":
		return "This field must be a boolean value"
	case "uuid", "uuid4":
		return "Invalid UUID format"
	case "datetime":
		return fmt.Sprintf("Invalid date format, must match: %s", fe.Param())
	default:
		return fmt.Sprintf("Validation failed on the '%s' rule", fe.Tag())
	}
}
