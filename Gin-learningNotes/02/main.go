package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 配置路由

	/* 四个方法 （的应对） */
	//1. GET：从服务器取出资源
	r.GET("/getwangzhi", func(c *gin.Context) {
		c.String(200, "Get")
	})
	//2. POST：在服务器新建资源
	r.POST("/postwangzhi", func(c *gin.Context) {
		c.String(200, "POST")
	})
	//3. PUT： 在服务器跟新资源（客户端提供改变后的完整资源）
	r.PUT("/putwangzhi", func(c *gin.Context) {
		c.String(200, "PUT")
	})
	//4. DELETE: 从服务器删除资源
	r.DELETE("网址", func(c *gin.Context) {
		c.String(200, "DELETE")
	})

	/*路由里面获取 Get 传值*/

	/*动态路由*/

	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run() //在9090端口使用：r.Run(":9090")
}
