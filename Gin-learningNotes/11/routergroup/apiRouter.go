package routergroup

import "github.com/gin-gonic/gin"

func ApiRouterInit(r *gin.Engine) {
	ApiRouter := r.Group("api")
	{
		ApiRouter.GET("/V1", func(c *gin.Context) {
			c.HTML(200, "default/default.html", gin.H{})
		})
		ApiRouter.GET("/V2", func(c *gin.Context) {
			c.HTML(200, "default/default.html", gin.H{})
		})
	}

}
