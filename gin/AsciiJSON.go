package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/someJSON", func(ctx *gin.Context) {

		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		ctx.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8080")
}
