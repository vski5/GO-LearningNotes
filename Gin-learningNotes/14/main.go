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
		username := c.PostForm("username")
		facefile, _ := c.FormFile("face")
		dst := path.Join("./default/", facefile.Filename)
		c.SaveUploadedFile(facefile, dst)
		//展示一下
		c.JSON(200, gin.H{
			"dst":      dst,
			"username": username,
		})
	})

	/*多个文件上传，不历遍的版本*/
	r.GET("/upload2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default/up2.html", gin.H{})

	})
	r.POST("/uploadpage2", func(c *gin.Context) {
		facefile, _ := c.FormFile("face")
		dst := path.Join("./default/", facefile.Filename)
		c.SaveUploadedFile(facefile, dst)
		//重复一遍就行
		facefile2, _ := c.FormFile("face2")
		dst2 := path.Join("./default/", facefile2.Filename)
		c.SaveUploadedFile(facefile, dst2)
	})

	/*多个文件上传，历遍的版本*/

	r.Run(":8080")

}
