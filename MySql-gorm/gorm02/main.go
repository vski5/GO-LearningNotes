package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct { //此时，默认的对应的表为users
	Id       int //都是表里的列，首字母大写。
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 就算不是英语也一样，Xxx这个结构体对应数据库里的xxxs，但也有例外
/* 在结构体上绑定方法 ， 改变结构体的默认表名称，此时User绑定的就从users变成了user，如果没有，那就查不到
func (U User) TableName() string {
	return "user"
} */

type Ceshi struct { //此时，默认的对应的表为ceshis <--小写末尾加s
	Id   int //都是表里的列，首字母大写。
	Punk int
}

func main() {
	r := gin.Default()

	// 连接数据库："user:password@tcp(127.0.0.1:3306)/数据库的名字?charset=utf8mb4&parseTime=True&loc=Local"
	//username:password@protocol(address)/dbname?param=value
	dsn := "root:password@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 改gorm.Config这个struct进行设置
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DB == nil {
		panic(1)
	}

	r.GET("/users", func(c *gin.Context) {

		//先实例化,把models.User{}先变成结构体对象的切片，方便models.DB.First(&user)等查询方法赋值
		userArry := []User{}

		// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
		DB.Find(&userArry)

		//测试一下
		c.JSON(200, gin.H{
			"userArry": userArry,
		})
	})
	r.GET("/ceshis", func(c *gin.Context) {

		//先实例化,把models.Ceshi{}先变成结构体对象的切片，方便赋值
		ceshisArry := []Ceshi{}

		// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
		DB.Find(&ceshisArry)

		//测试一下
		c.JSON(200, gin.H{
			"ceshisArry": ceshisArry,
		})
	})

	r.Run(":8080")
}
