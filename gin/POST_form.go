package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/save", func(ctx *gin.Context) {
		name := ctx.Query("name")
		age := ctx.DefaultQuery("age", "0")
		form_msg := ctx.PostForm("form_msg")

		fmt.Printf("save name:%s, age:%s, form_msg:%s", name, age, form_msg)

		ctx.AsciiJSON(http.StatusOK, "save success")
	})

	r.Run(":8080")
}
