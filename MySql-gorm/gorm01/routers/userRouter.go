package routers

import (
	"gorm01/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouterInit(r *gin.Engine) {
	userRouters := r.Group("/db")
	{
		userRouters.GET("/add", controllers.UserController{}.Add)
		userRouters.GET("/search", controllers.UserController{}.Search)
		userRouters.GET("/delete", controllers.UserController{}.Delete)
	}

}
