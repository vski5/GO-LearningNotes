package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("different/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/index.html", gin.H{
			"title": "标题",
			"real": &Article{
				Title:   "test1",
				Content: "test2",
			},
		})
	})

	r.Run(":8080")
}
