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
	r.POST("/adduser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20000") //在HTML里写了表单，即使不填也会发送“空”进行填空，不会出现第二个空设置的默认值

		c.JSON(200, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})

	r.Run(":8080")
}
