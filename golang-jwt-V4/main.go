package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/golang-jwt/jwt/v4"
)

// 这里定义JWT的Payload部分和加密密码和作用域
type MyUserJWTClaims struct {
	//Id       int      `json:"id"`
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

// 生成和签名JWT
func (userJWT MyUserJWTClaims) CreateToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{ //MapClaims is a claims type that uses the map[string]interface{} for JSON decoding.
			//"Id":       userJWT.Id,
			"Username": userJWT.Username,
			"Urls":     userJWT.Urls,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString(userJWT.Secret)

	return tokenString, err

}

// 验证 和 解析 JWT
func ValidateToken(tokenString string, my_secret_key []byte) (back map[string]interface{}, err error) {

	if tokenString == " " { //如果没有Token

		//c.String(http.StatusInternalServerError, err.Error())
		return nil, errors.New("token error")
	} else { //token和 JWT 装饰器函数传递到 parse 方法中，然后返回接口和错误
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
				//id := claims["Id"]
				username := claims["Username"]
				urls := claims["Urls"]

				userJWTS := map[string]interface{}{
					//"Id":       id,
					"Username": username,
					"Urls":     urls,
				}
				return userJWTS, errors.New("token error")
			}

		}

	}
	return nil, errors.New("validate token error")

}

// 将使用GIN设置JWT 和 生成和签名JWT 封装到一起。 Scope是 作用域
func (userJWT MyUserJWTClaims) SetJWT(c *gin.Context, Scope string) {
	tokenString, err := userJWT.CreateToken()
	if tokenString != "" && err == nil {
		c.SetCookie("Token", tokenString, 3600, "/", Scope, false, true)
	}
}

// 将使用GIN获取JWT 和 验证并且解析JWT 封装到一起。需要实例化MyUserJWTClaims里的Secret
func (userJWT MyUserJWTClaims) GetJWT(c *gin.Context) ([]string, error) {
	// 从请求头中获取 Token 信息
	tokenget111, _ := c.Request.Cookie("Token")

	ca, _ := ValidateToken(tokenget111.Value, userJWT.Secret)

	if urls, ok := ca["Urls"].([]interface{}); ok {
		y := make([]string, len(urls))
		for i, v := range urls {
			y[i] = v.(string)
		}
		// y 是 []string 类型
		return y, nil

	} else {
		// x["Urls"] 不是 []interface{} 类型
		return nil, errors.New("x[\"Urls\"] 不是 []interface{} 类型")
	}
}

// 在gin环境下测试
func main() {
	r := gin.Default()

	r.GET("/set", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Username: "JWTUsernameJWTUsernameJWTUsernameJWTUsername",
			Urls:     []string{"test213456798", "test213456798"},
			Secret:   []byte("my_secret_key"),
		}
		val, _ := a.CreateToken()

		// 将token存储到客户端的浏览器中
		//c.Set("Token", val)
		c.SetCookie("Token", val, 3600, "/", "a.20011111.xyz", false, true)
		c.String(200, val)
	})

	r.GET("/get", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Secret: []byte("my_secret_key"),
		}
		// 从请求头中获取 Token 信息
		tokenget111, _ := c.Request.Cookie("Token")

		ca, _ := ValidateToken(tokenget111.Value, a.Secret)

		if urls, ok := ca["Urls"].([]interface{}); ok {
			y := make([]string, len(urls))
			for i, v := range urls {
				y[i] = v.(string)
			}
			// y 是 []string 类型
			fmt.Println(y)
			//我需要得到的就是y
			fmt.Println("===================我需要得到的就是 y ======================")
		} else {
			// x["Urls"] 不是 []interface{} 类型
			fmt.Println("x[\"Urls\"] 不是 []interface{} 类型")
			fmt.Println("=========================================")
		}

	})
	r.GET("/test1", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Username: "JWTUsernameJWTUsernameJWTUsernameJWTUsername",
			Urls:     []string{"test213456798", "test213456798"},
			Secret:   []byte("my_secret_key"),
		}
		a.SetJWT(c, "a.20011111.xyz")
	})
	r.GET("/test2", func(c *gin.Context) {
		a := &MyUserJWTClaims{
			Username: "JWTUsernameJWTUsernameJWTUsernameJWTUsername",
			Urls:     []string{"test213456798", "test213456798"},
			Secret:   []byte("my_secret_key"),
		}
		str, _ := a.GetJWT(c)
		fmt.Println(str)
		//用反射求证一下
		t := reflect.TypeOf(str)
		fmt.Println(t.Name())
		fmt.Println(t.Kind())
	})
	r.Run(":8080")

}
