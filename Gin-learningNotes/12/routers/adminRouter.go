package routers

import (
	"12/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("admin")
	{
		adminRouter.GET("test1", admin.AdminController.Init)
		adminRouter.GET("test2", func(c *gin.Context) {
			c.String(200, "test2", gin.H{})
		})

	}
}
