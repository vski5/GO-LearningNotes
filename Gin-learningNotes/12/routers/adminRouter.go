package routers

import (
	"12/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouter := r.Group("admin")
	{
		adminRouter.GET("/test1", admin.AdminController{}.Add)
		adminRouter.GET("/test2", admin.AdminController{}.Back)

	}
}
