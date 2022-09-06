package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Artitle struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	//r.LoadHTMLFiles("default/index.html", "default/news.html") 因为r.LoadHTMLFiles不支持泛解析
	//r.LoadHTMLGlob("default/*")可以直接准备一个文件夹
	r.LoadHTMLGlob("different/**/*") //有两个同名文件在不同文件夹下面的解决方法，记得在HTML首位加上{{ define "文件夹名字/文件名字" }}和{{<end>}}
	r.GET("/html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html",
			gin.H{
				"title": "text",
			})

	})
	r.GET("/html2", func(c *gin.Context) {
		c.HTML(http.StatusOK, "default1/index.html",
			gin.H{
				"title": "text2",
			})

	})

	//嵌套结构的html渲染

	r.GET("/html3", func(c *gin.Context) {
		//实例化一下
		news := &Artitle{
			Title:   "标题",
			Content: "内容",
		}
		c.HTML(200, "default1/news.html",
			gin.H{
				"title": "主标题",
				"news":  news,
			})
	})

	r.Run(":8080")
}
