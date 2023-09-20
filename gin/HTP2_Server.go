package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var html = template.Must(template.New("https").Parse(`
<html>
<head>
  <title>Https Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	r := gin.Default()
	r.Static("assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(ctx *gin.Context) {
		if pusher := ctx.Writer.Pusher(); pusher != nil {
			//使用pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		ctx.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})
	// 监听并在 https://127.0.0.1:8080 上启动服务
	r.RunTLS(":8080", "./testdata/server/pem", "./testdata/server.key")
}
