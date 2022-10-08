package middleware

import (
	"fmt"

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
	fmt.Println("一个-路由组-中间件")
}
