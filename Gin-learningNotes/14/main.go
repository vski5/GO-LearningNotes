// Gin中实现单文件上传 多文件上传
package main

import (
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("static", "./static")
	r.LoadHTMLGlob("./default/*")

	//单个文件上传，用到c.FormFile获得文件，c.SaveUploadedFile储存文件，path.join拼接储存地址。
	//第一步，提供页面
	r.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/up.html", gin.H{})

	})
	//第二步，路径第一步的HTML中form路径相同，类似一个Hook绑在第一步。
	r.POST("/uploadpage", func(c *gin.Context) {
		facefile, _ := c.FormFile("face")
		dst := path.Join("./default/", facefile.Filename)
		c.SaveUploadedFile(facefile, dst)
	})

	r.Run(":8080")

}
