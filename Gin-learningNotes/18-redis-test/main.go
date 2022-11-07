package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rueian/rueidis"
)

var Rueidisctx context.Context
var RedisStore rueidis.Client
var err error

func init() {
	//连接redis
	RedisStore, err = rueidis.NewClient(rueidis.ClientOption{
		//写链接
		InitAddress: []string{"127.0.0.1:6379"},
		//写密码
		Password: "mimaisc1r89ev5",
	})

	//检错
	if err != nil {
		panic(err)
	}
	//结束时的关闭
	//defer redisStore.Close()

	//redis共享的上下文
	Rueidisctx = context.Background()
}
func main() {
	//设置gin
	r := gin.Default()

	// SET key val NX，设置key/value
	//redisStore.Do(rueidisctx, redisStore.B().Set().Key("key").Value("val").Nx().Build()).Error()
	// GET key
	//redisStore.Do(rueidisctx, redisStore.B().Get().Key("key").Build()).ToString()
	r.GET("redischeck", func(c *gin.Context) {
		//设置key/value
		RedisStore.Do(Rueidisctx, RedisStore.B().Set().Key("key111").Value("val999").Nx().Build()).Error()

		//GET key
		username, _ := RedisStore.Do(Rueidisctx, RedisStore.B().Get().Key("key111").Build()).ToString()

		c.JSON(200, gin.H{
			"放redis里的cookie的key":   "userkey",
			"放redis里的cookie的value": username,
		})
	})
	r.Run(":8080")
}
