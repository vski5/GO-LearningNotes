package admin

import "github.com/gin-gonic/gin"

type AdminController struct{}

func (a *AdminController) Add(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{})
}
func (a *AdminController) BackString(c *gin.Context) {
	c.String(200, "test2", gin.H{})
}
