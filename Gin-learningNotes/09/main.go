package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("static", "./static")
	r.LoadHTMLGlob("formpages/*")
	r.GET("/", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		page := c.DefaultQuery("page", "no page") //第二个是没传参情况的默认值
		//传参网址为：/?name=nametest&age=agetest  ,用?表示要传参了，两个参数之间用&
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"page": page,
		})
	})

	//POST传递表单的方法
	//第一步：
	r.GET("/forms", func(c *gin.Context) {
		c.HTML(200, "formpages/formpage.html", gin.H{})
	})
	//第二部：
	r.POST("/addform", func(c *gin.Context) {
		zhanghao := c.PostForm("zhanghao")
		mima := c.PostForm("mima")

		c.JSON(200, gin.H{
			"zhanghao": zhanghao,
			"mima":     mima,
		})
	})

	r.Run(":8080")
}
