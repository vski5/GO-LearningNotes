//路由中的方法分离
package main

import (
	"12/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("html/**/*")
	routers.AdminRouterInit(r)
	routers.FrontRouterInit(r)
	r.Run(":8080")
}
