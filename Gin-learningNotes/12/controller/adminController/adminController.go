package admincontroller

import "github.com/gin-gonic/gin"

type AdminController struct{}

func (a *AdminController) Init(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{})
}
