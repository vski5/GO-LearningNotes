package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//本质上就是自己建一个包里面有函数，但是更加有条理,这些函数写在--路由组分组--文件里面。
func AdminRoutergroupMiddleware1(c *gin.Context) {

	fmt.Println("一个-路由组-中间件")
}
func AdminRoutergroupMiddleware2(c *gin.Context) {
	//用Get获取在13/controller/front/frontController.go里面设置的c.Set
	usernameget, _ := c.Get("username")
	v, ok := usernameget.(string) //类型断言
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}

	//gin 中间件中使用 goroutine不能使用原始的上下文（c *gin.Context），必须使用其只读副本 c.Copy()
	cCopy := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		// 这里使用你创建的副本
		fmt.Println("Done! in path " + cCopy.Request.URL.Path)

	}()
	fmt.Println("一个-路由组-中间件")
}
