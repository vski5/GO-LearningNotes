package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	/*将session保存在redis里面*/
	// 第 1 个参数 - redis 最大的空闲连接数
	// 第 2 个参数 - 数通信协议 tcp 或者 udp
	// 第 3 个参数 - redis 地址, 格式，host:port
	// 第 4 个参数 - redis 密码
	// 第 5 个参数 - session 加密密钥
	redisStore, _ := redis.NewStore(5, "tcp", "127.0.0.1:6379", "", []byte("sessionKeyWord"))

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
		username := session.Get("rediscookie999")
		c.JSON(200, gin.H{
			"放redis里的cookie的key":   "rediscookie999",
			"放redis里的cookie的value": username,
		})
	})
	r.Run(":8080")
}
