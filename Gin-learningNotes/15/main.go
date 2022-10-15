// Gin中实现 cookie（本地）和session（客户端）
package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

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

	/*搭建 Session*/

	//存储在基于cookie的存储引擎
	//创建基于cookie的存储引擎,[]byte("keyPassword")是密码
	cookieSaver := cookie.NewStore([]byte("keyPassword"))
	//package sessions ("github.com/gin-contrib/sessions")
	//设置全局中间件（所有路由都调用这个函数）
	//传入值：cookie名，保存在哪个储存引擎
	r.Use(sessions.Sessions("cookie_session_name111", cookieSaver))
	r.GET("/sessionSet", func(c *gin.Context) {
		//初始化session,表明与gin公用一个上下文*gin.Context
		session111 := sessions.Default(c)
		//通过修改struct也就是sessions.Options的内容来设置session存活时间
		session111.Options(sessions.Options{
			MaxAge: 3600 * 3, //三小时
		})
		//把session111同步设置到全局
		session111.Set("cookieName222", "cookieValue222")
		session111.Save()
		//返回一个json看看效果
		c.JSON(200, gin.H{
			"获取cookie的对应value": session111.Get("cookieName222"),
		})
	})
	r.GET("sessionCheck", func(c *gin.Context) {
		//初始化session,只要是共用gin这个上下文就是一致的
		session999 := sessions.Default(c)
		c.JSON(200, gin.H{
			"获取cookie的对应value": session999.Get("cookieName222"),
		})
	})

	/*将session保存在redis里面*/
	// 第 1 个参数 - redis 最大的空闲连接数
	// 第 2 个参数 - 数通信协议 tcp 或者 udp
	// 第 3 个参数 - redis 地址, 格式，host:port
	// 第 4 个参数 - redis 密码
	// 第 5 个参数 - session 加密密钥
	redisStore, _ := redis.NewStore(5, "tcp", "localhost:6379", "", []byte("sessionKeyWord"))

	//设置中间件到全局
	r.Use(sessions.Sessions("mySessionName", redisStore))
	r.GET("redis", func(c *gin.Context) {
		redissession := sessions.Default(c)
		redissession.Set("rediscookie", "redisname")
		redissession.Save()
		c.JSON(200, map[string]interface{}{
			"放redis里的cookie的key":   "rediscookie",
			"放redis里的cookie的value": redissession.Get("rediscookie"),
		})
	})
	//先在redis-cli里set rediscookie222 redisvalue222
	//再确认
	r.GET("redischeck", func(c *gin.Context) {
		// 初始化 session 对象
		session := sessions.Default(c)
		// 通过 session.Get 读取 session 值
		username := session.Get("rediscookie222")
		c.JSON(200, gin.H{
			"放redis里的cookie的key":   "rediscookie222",
			"放redis里的cookie的value": username,
		})
	})

	r.Run(":8080")
}
