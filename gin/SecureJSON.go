package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/secureJSON", func(ctx *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		ctx.SecureJSON(http.StatusOK, names)
	})

	r.Run(":8080")
}
