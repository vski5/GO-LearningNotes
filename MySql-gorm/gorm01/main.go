package main

import (
	"gorm01/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userRouters := r.Group("admin")
	{
		userRouters.GET("/add", controllers.UserController{}.Add)
		userRouters.GET("/search", controllers.UserController{}.Search)
	}
	r.Run(":8080")
}
