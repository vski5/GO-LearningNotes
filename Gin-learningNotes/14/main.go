//Gin中实现单文件上传 多文件上传
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET()
	r.Run(":8080")
}
