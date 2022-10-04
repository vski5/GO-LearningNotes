package routers

import "github.com/gin-gonic/gin"

func FrontRouterInit(c *gin.Engine) {
	frontRouter := c.Group("front")
	{
		frontRouter.GET("/test1", func(c *gin.Context) {
			c.HTML(200, "html2/html2.html", gin.H{})
		})
		frontRouter.GET("/test2", func(c *gin.Context) {
			c.String(200, "test222", gin.H{})
		})
	}
}
