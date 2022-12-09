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
	Url      []string
	Secret   string
	//Scope    string //作用域
}

// 生成JWT，并且用gin存储到浏览器中
func (userJWT MyUserJWTClaims) CreateToken() (string, error) {
	//用HS256生成一个新的JWT
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	//设置过期时间为24小时
	exp := time.Now().Add(time.Hour * 24).Unix()
	// 添加JWT 中包含的声明（claims）
	//jwt.MapClaims 是 Go 语言中的一个类型，它表示 JWT 中的有效负载
	claims := jwt.MapClaims{
		"Id":       userJWT.Id,
		"Username": userJWT.Username,
		"Url":      userJWT.Url,
		"exp":      exp,
	}
	token.Claims = claims
	// 用密码签名 JWT，并返回签名后的 JWT 字符串
	val, err := token.SignedString(userJWT.Secret)

	if err != nil {
		fmt.Println(" token.SignedString __ error", err)
		return " ", err
	}

	return val, err

	// 将token存储到客户端的浏览器中
	//c.SetCookie("token", val, 3600, "/", "127.0.0.1", false, true)

}

// 在gin环境下测试
func main() {
	r := gin.Default()

	r.GET("/test111", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Id:       1,
			Username: "JWTUsernameJWTUsernameJWTUsernameJWTUsername",
			Url:      []string{"test213456798", "test213456798"},
			Secret:   "11111",
		}
		val, _ := a.CreateToken()
		// 将token存储到客户端的浏览器中
		//c.SetCookie("token", val, 3600, "/", "127.0.0.1", false, true)
		c.String(200, val)
	})

	r.Run(":8080")

}
