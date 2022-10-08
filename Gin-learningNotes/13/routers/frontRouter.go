package routers

import (
	"13/controller/front"
	"fmt"

	"github.com/gin-gonic/gin"
)

func useMiddleware(c *gin.Context) {
	fmt.Println("一个-路由组-中间件")
}
func useMiddleware2(c *gin.Context) {
	fmt.Println("一个-路由组-中间件")
}
func FrontRouterInit(c *gin.Engine) {
	frontRouter := c.Group("front", useMiddleware) //第二个及之后的都是路由组中间件
	//也可以写在下面
	frontRouter.Use(useMiddleware, useMiddleware2)
	{
		frontRouter.GET("/test1", front.FrontController{}.BackHTML)
		frontRouter.GET("/test2", front.FrontController{}.BackString)

		frontRouter.GET("/test3", front.FrontController{}.Test)
	}
}
