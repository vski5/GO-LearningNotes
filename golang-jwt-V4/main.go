package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v4"
)

// 这里定义JWT的Payload部分和加密密码和作用域
type MyUserJWTClaims struct {
	Id       int      `json:"id"`
	Username string   `json:"username"`
	Urls     []string `json:"urls"`
	Secret   []byte   //[]byte("my_secret_key")
	//Scope    string   //作用域
	/*
			jwt.RegisteredClaims这个struct用于存储已注册的声明（claims），并包含了一些公用的声明。

		Issuer字段存储了声明的发行者（issuer）。
		Subject字段存储了声明的主题（subject）。
		Audience字段存储了声明的接收方（audience）。type ClaimStrings []string
		ExpiresAt字段存储了声明的过期时间（expiration time）.*NumericDate。type NumericDate struct { time.Time }
		NotBefore字段存储了声明的生效时间（not before）。 type NumericDate struct { time.Time }
		IssuedAt字段存储了声明的签发时间（issued at）。 type NumericDate struct { time.Time }
		ID字段存储了声明的唯一标识符（JWT ID）。
	*/
}

// 生成和签名JWT，并且用gin存储到浏览器中
func (userJWT MyUserJWTClaims) CreateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{ //MapClaims is a claims type that uses the map[string]interface{} for JSON decoding.
			"Id":       userJWT.Id,
			"Username": userJWT.Username,
			"Urls":     userJWT.Urls,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(userJWT.Secret)

	return tokenString, err

}

// 验证 和 解析 JWT
func ValidateToken(c *gin.Context, my_secret_key []byte) (back interface{}, err error) {
	// 获取 http.Request 类型的指针
	req := c.Request

	// 从请求头中获取 Token 信息
	tokenget := req.Header.Get("Token")

	if tokenget == " " { //如果没有Token

		//c.String(http.StatusInternalServerError, err.Error())
		return fmt.Errorf("token is nill"), nil
	} else { //token和 JWT 装饰器函数传递到 parse 方法中，然后返回接口和错误
		token, _ := jwt.Parse(tokenget, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return fmt.Errorf("there was an error in parsing"), nil //错误消息首字母要小写。创建一个 error错误。
			}
			return my_secret_key, nil //[]byte("my_secret_key")记得改成自己的密码，用struct解决。
		})
		if token == nil {
			fmt.Println("invalid token")
		} else {
			// 解析JWT
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				fmt.Println("couldn't parse claims")
				return nil, errors.New("token error")
			} else {
				// 以MyUserJWTClaims构建的JWT为例

				id := claims["Id"].(int)
				username := claims["Username"].(string)
				urls := claims["Urls"].([]string)
				userJWTS := &MyUserJWTClaims{
					Id:       id,
					Username: username,
					Urls:     urls,
				}

				return userJWTS, errors.New("token error")
			}

		}

	}
	return nil, errors.New("validate token error")

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
		ca, _ := ValidateToken(c, a.Secret)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		//fmt.Println(a1)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		//fmt.Println(b1)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(ca)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(val, err111)
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		// 将token存储到客户端的浏览器中
		//c.SetCookie("token", val, 3600, "/", "localhost", false, true)
		c.String(200, val)
	})

	r.Run(":8080")

}
