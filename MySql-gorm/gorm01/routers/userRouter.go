package routers

import (
	"gorm01/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("/db")
	{
		//增
		userRouters.GET("/add", controllers.UserController{}.Add)
		userRouters.GET("/add2", controllers.UserController{}.Add2)
		//查
		userRouters.GET("/search", controllers.UserController{}.Search)
		userRouters.GET("/search2", controllers.UserController{}.Search2)
		userRouters.GET("/search3", controllers.UserController{}.Search3)
		//删
		userRouters.GET("/delete", controllers.UserController{}.Delete)
		//更新
		userRouters.GET("/update", controllers.UserController{}.Update1)
	}

}
