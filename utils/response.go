package utils

import (
	"github.com/gin-gonic/gin"
)

// Response adalah format standar untuk semua API
type Response struct {
	Status  string `json:"status"`           // "success" atau "fail" / "error"
	Message string `json:"message"`          // Pesan singkat penjelasan
	Data    any    `json:"data,omitempty"`   // Data utama (bisa null/empty)
	Errors  any    `json:"errors,omitempty"` // Khusus untuk validation errors
}

// SuccessResponse mengirim response 200/201 dengan format seragam
func SuccessResponse(c *gin.Context, code int, message string, data any) {
	c.JSON(code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// ErrorResponse mengirim response error (4xx/5xx) dengan format seragam
func ErrorResponse(c *gin.Context, code int, message string, errors any) {
	c.JSON(code, Response{
		Status:  "fail",
		Message: message,
		Errors:  errors,
	})
}
