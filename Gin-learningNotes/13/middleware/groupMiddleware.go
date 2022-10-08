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
	fmt.Println("一个-路由组-中间件")
}
