package admin

import "github.com/gin-gonic/gin"

type AdminController struct{}

func Add(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{})
}
func Back(c *gin.Context) {
	c.String(200, "test2")
}
