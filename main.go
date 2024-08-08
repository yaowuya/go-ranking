package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	err := r.Run(":9999")
	if err != nil {
		return
	}

}
