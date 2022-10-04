package routers

import "github.com/gin-gonic/gin"

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("admin")
	{
		adminRouter.GET("test1", func(c *gin.Context) {
			c.HTML(200, "html1/html1.html", gin.H{})
		})
		adminRouter.GET("test2", func(c *gin.Context) {
			c.String(200, "test2", gin.H{})
		})

	}
}
