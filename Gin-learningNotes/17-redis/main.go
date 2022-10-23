//以后再探索redis的session储存
//用rueidis来实现redis的session储存

package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/rueian/rueidis"
)

func main() {
	//设置gin
	r := gin.Default()
	//连接redis
	redisStore, err := rueidis.NewClient(rueidis.ClientOption{
		//写链接
		InitAddress: []string{"127.0.0.1:6379"},
		//写密码
		Password: "",
	})

	//检错
	if err != nil {
		panic(err)
	}
	//结束时的关闭
	//defer redisStore.Close()

	//redis共享的上下文
	rueidisctx := context.Background()

	// SET key val NX，设置key/value
	//redisStore.Do(rueidisctx, redisStore.B().Set().Key("key").Value("val").Nx().Build()).Error()
	// GET key
	//redisStore.Do(rueidisctx, redisStore.B().Get().Key("key").Build()).ToString()
	r.GET("redischeck", func(c *gin.Context) {
		//设置key/value
		redisStore.Do(rueidisctx, redisStore.B().Set().Key("key111").Value("val111").Nx().Build()).Error()

		//GET key
		username, _ := redisStore.Do(rueidisctx, redisStore.B().Get().Key("key111").Build()).ToString()

		c.JSON(200, gin.H{
			"放redis里的cookie的key":   "userkey",
			"放redis里的cookie的value": username,
		})
	})
	r.Run(":8080")
}
