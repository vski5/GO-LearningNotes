package front

import "github.com/gin-gonic/gin"

type FrontController struct{}

func BackHTML(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{})
}
func BackString(c *gin.Context) {
	c.String(200, "test2", gin.H{})
}
func (f FrontController) Test(c *gin.Context) {
	c.String(200, "test,test,test", gin.H{})
}
