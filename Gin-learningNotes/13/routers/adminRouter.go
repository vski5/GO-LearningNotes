package routers

import (
	"13/controller/admin"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	adminRouters := r.Group("admin")
	{
		//一个路由可以配多个回调函数，每一个就是中间件(HOOK)
		adminRouters.GET("/test1", admin.Middleware{}.Middleware1, admin.Middleware{}.Middleware2, admin.AdminController{}.Add)
		adminRouters.GET("/test2", admin.AdminController{}.Back)
	}
}
