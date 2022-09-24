package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userinfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"password"`
	Age      string `from:"age" json:"age"`
}

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
	//第三步，把传回的数据绑定到结构体,另用一个html页面。
	r.GET("/forms2", func(c *gin.Context) {
		c.HTML(200, "formpages/formpage2.html", gin.H{})
	})
	r.POST("/adduser2", func(c *gin.Context) {
		user := &Userinfo{}
		//理论上c.ShouldBind(&user)可以替代下面的三行c.PostForm，但是实操发现会失去age的值
		/* 		user.Username = c.PostForm("username")
		   		user.Password = c.PostForm("password")
		   		user.Age = c.PostForm("age") */ //在HTML里写了表单，即使不填也会发送“空”进行填空，不会出现第二个空设置的默认值
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, user)
			fmt.Printf("%#v", user)
		}
	})

	r.Run(":8080")
}
