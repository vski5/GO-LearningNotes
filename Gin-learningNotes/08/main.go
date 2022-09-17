//08：静态文件服务
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//"./static"是实际地址， "/statictest"是映射地址，浏览器求的地址是映射地址，GET "/statictest/css/style.css"
	r.Static("/statictest", "./static")
	r.LoadHTMLGlob("different/**/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/index.html", gin.H{
			"title": "标题",
		})
	})

	r.Run(":8080")
}
