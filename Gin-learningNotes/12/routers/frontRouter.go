package routers

import (
	"12/controller/front"

	"github.com/gin-gonic/gin"
)

func FrontRouterInit(c *gin.Engine) {
	frontRouter := c.Group("front")
	{
		frontRouter.GET("/test1", front.BackHTML)
		frontRouter.GET("/test2", front.BackString)

		frontRouter.GET("/test3", front.FrontController{}.Test)
	}
}
