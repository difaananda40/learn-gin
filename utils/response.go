package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func SuccessResponse(c *gin.Context, code int, message string, data any) {
	c.JSON(code, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, message string, errors any) {
	c.JSON(code, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
