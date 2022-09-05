package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("default/index.html", "default/news.html") 因为r.LoadHTMLFiles不支持泛解析
	//r.LoadHTMLGlob("default/*")可以直接准备一个文件夹
	r.LoadHTMLGlob("default/*")
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html",
			gin.H{
				"title": "text",
			})
	})
	r.Run(":8080")
}
