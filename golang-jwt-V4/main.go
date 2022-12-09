package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

// 这里定义JWT的Payload部分和加密密码和作用域
type MyUserJWTClaims struct {
	Id       int
	Username string
	Urls     []string
	Secret   []byte //[]byte("my_secret_key")
	Scope    string //作用域
}

// 生成JWT，并且用gin存储到浏览器中
func (userJWT MyUserJWTClaims) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Id":       userJWT.Id,
		"Username": userJWT.Username,
		"Url":      userJWT.Urls,
		"endTime":  time.Now().Add(time.Hour * 24),
	})
	tokenString, err := token.SignedString(userJWT.Secret)
	return tokenString, err

}

// 在gin环境下测试
func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Id:       1,
			Username: "JWTUsernameJWTUsernameJWTUsernameJWTUsername",
			Urls:     []string{"test213456798", "test213456798"},
			Secret:   []byte("my_secret_key"),
		}
		val, err111 := a.CreateToken()
		fmt.Println(val, err111)
		// 将token存储到客户端的浏览器中
		//c.SetCookie("token", val, 3600, "/", "a.20011111.xyz", false, true)
		c.String(200, val)
	})

	r.Run(":8080")

}
