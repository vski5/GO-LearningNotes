package admin

import "github.com/gin-gonic/gin"

type AdminController struct{}

func (a AdminController) Add(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{"title": "标题", "real": "AdminController"})
}
func (a AdminController) Back(c *gin.Context) {
	c.String(200, "test2")
}
