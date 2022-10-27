package main

import (
	_ "gorm01/models"
	"gorm01/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.UserRouterInit(r)
	r.Run(":8080")

}
