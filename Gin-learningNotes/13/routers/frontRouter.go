package routers

import (
	"13/controller/front"

	"github.com/gin-gonic/gin"
)

func FrontRouterInit(c *gin.Engine) {
	frontRouter := c.Group("front")
	{
		frontRouter.GET("/test1", front.FrontController{}.BackHTML)
		frontRouter.GET("/test2", front.FrontController{}.BackString)

		frontRouter.GET("/test3", front.FrontController{}.Test)
	}
}
