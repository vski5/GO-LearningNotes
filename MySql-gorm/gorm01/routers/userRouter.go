/* package routers

import (
	"gorm01/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouterInit() (rou *gin.Engine) {
	userRouters := rou.Group("admin")
	{
		userRouters.GET("/add", controllers.UserController{}.Add)
		userRouters.GET("/search", controllers.UserController{}.Search)
	}
}
*/