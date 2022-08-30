package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("running")
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 配置路由
	//对着，用get方法访问了后缀为check网址的用户，使用后面的func函数
	r.GET("/check", func(c *gin.Context) {
		c.JSON(200, gin.H{ // c.JSON：返回 JSON 格式的数据
			"message": "Hello world!"})
	})
	// 启动 HTTP 服务，默认在 0.0.0.0:8080 启动服务
	r.Run() //在9090端口使用：r.Run(":9090")
}

//热加载：在命令行输入 go run github.com/pilu/fresh
//感觉比Air简单
