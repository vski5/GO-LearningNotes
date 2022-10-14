// Gin中实现 cookie（本地）和session（客户端）
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/pathtest/home", func(c *gin.Context) {
		//设置cookie，第5个参数表示所有20011111.xyz的二级域名（比如a.20011111.xyz b.20011111.xyz）都共享一个cookie
		//第4个表示路径，我猜测是所有/pathtest路径的共享这个cookie，如果设置为/，就是所有都共享。
		c.SetCookie("cookiename", "cookieValue", 360, "/pathtest", ".20011111.xyz", false, false)
		//获取cookie，这里是仅用于展示
		cookie111, _ := c.Cookie("cookiename")
		c.String(200, "这是设置cookie1的默认界面,本次设置的cookie名字是《%v》", cookie111)
	})
	r.GET("/pathtest/pathtest", func(c *gin.Context) {
		//获取cookie，这里是仅用于展示
		cookie111, _ := c.Cookie("cookiename")
		c.String(200, "这是设置cookie1的默认界面,本次设置的cookie名字是《%v》", cookie111)
	})
	r.GET("/home", func(c *gin.Context) {
		//获取cookie，这里是仅用于展示
		cookie111, _ := c.Cookie("cookiename")
		c.String(200, "测试cookie《%v》", cookie111)
	})

	r.GET("/delete", func(c *gin.Context) {
		//删除cookie，第三个参数是cookie存在的时间，复数就是删去，可以把这个函数绑在一个Hook后面达到删cookie的目的，
		//此处是访问这个页面就删名为cookiename的cookie
		//第四个表示删除/后面的所有cookie
		c.SetCookie("cookiename", "cookieValue", -1, "/", ".20011111.xyz", false, false)

		c.String(200, "这是删除cookie")
	})

	r.Run(":8080")
}
