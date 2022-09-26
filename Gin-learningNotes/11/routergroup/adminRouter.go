package routergroup

import "github.com/gin-gonic/gin"

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("admin")
	{
		adminRouter.GET("/V1", func(c *gin.Context) {
			c.HTML(200, "admin/admin.html", gin.H{})
		})
		adminRouter.GET("/V2", func(c *gin.Context) {
			c.HTML(200, "admin/admin.html", gin.H{})
		})
	}

}
