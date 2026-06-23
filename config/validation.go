package config

import (
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func InitializeValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			jsonTag := fld.Tag.Get("json")
			name, _, _ := strings.Cut(jsonTag, ",")
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
