//路由中的方法分离
package main

import (
	"13/routers"
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

func SplicingFunc(input1, input2 string) string {
	return input1 + "-----" + input2 //类java的+字符拼凑
}
func useMiddleware(c *gin.Context) {
	fmt.Println("一个全局中间件")
}
func useMiddleware2(c *gin.Context) {
	fmt.Println("一个全局中间件222")
}
func main() {
	r := gin.Default()
	r.Use(useMiddleware, useMiddleware2) //Use制造全局中间件，所有路由后面都第一个执行一次。可以加无数个全局中间件
	r.Static("/static", "./static")
	r.SetFuncMap(template.FuncMap{
		"Splicing": SplicingFunc,
	})
	r.LoadHTMLGlob("html/**/*")

	routers.AdminRouterInit(r)
	routers.FrontRouterInit(r)
	r.Run(":8080")
}
