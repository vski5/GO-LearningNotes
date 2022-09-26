//路由分组
package main

import (
	"11/routergroup"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("html/**/*")
	routergroup.ApiRouterInit(r)
	routergroup.AdminRouterInit(r)
	r.Run(":8080")
}
