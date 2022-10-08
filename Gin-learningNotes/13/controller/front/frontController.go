package front

import (
	"github.com/gin-gonic/gin"
)

type FrontController struct{}

func (f FrontController) BackHTML(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{"title": "标题", "real": "FrontController"})
}
func (f FrontController) BackString(c *gin.Context) {
	c.String(200, "test2", gin.H{})
}
func (f FrontController) Test(c *gin.Context) {

	c.Set("username", "张三") //存储一个新的键/值对，专门用于这个上下文，可以被所有用了这个*gin.Context的get到

	c.String(200, "test,test,test", gin.H{})
}
