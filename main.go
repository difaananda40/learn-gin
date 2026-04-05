package main

import (
	"learn-gin/api/post"

	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	baseRoute := router.Group("/api")

	// Validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	post.RegisterHandlers(baseRoute)

	router.Run() // listens on 0.0.0.0:8080 by default
}
