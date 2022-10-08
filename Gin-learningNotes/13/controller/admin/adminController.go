package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AdminController struct{}

func (a AdminController) Add(c *gin.Context) {
	c.HTML(200, "html1/html1.html", gin.H{"title": "标题", "real": "AdminController"})
}
func (a AdminController) Back(c *gin.Context) {
	c.String(200, "test2")
}

//中间件
type Middleware struct{}

func (m Middleware) Middleware1(c *gin.Context) {

	fmt.Println("111")
	//c.Next()会直接跳过下面的函数，执行下一个回调函数（中间件），最后在从第一个回调函数的c.Next()后面的执行
	c.Next()
	fmt.Println("222")

}
func (m Middleware) Middleware2(c *gin.Context) {
	fmt.Println("AAAAA")

	//c.Abort()触发会直接终止路由里面所有控制器的后续进行，没有阻止此处中间件的输出，但是阻止了后面c.HTML这个中间件。
	c.Abort()

	fmt.Println("BBB")
}
