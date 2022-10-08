//路由中的方法分离
package main

import (
	"13/routers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func SplicingFunc(input1, input2 string) string {
	return input1 + "-----" + input2 //类java的+字符拼凑
}

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.SetFuncMap(template.FuncMap{
		"Splicing": SplicingFunc,
	})
	r.LoadHTMLGlob("html/**/*")
	routers.AdminRouterInit(r)
	routers.FrontRouterInit(r)
	r.Run(":8080")
}
