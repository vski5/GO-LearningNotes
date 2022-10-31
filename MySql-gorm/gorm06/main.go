package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 中使用事务
//初始化时禁用 GORM事务 ，这将获得大约 30%+ 性能提升
/*
gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	跳过默认设置
*/

func main() {
	r := gin.Default()

	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 改gorm.Config这个struct进行设置
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if DB == nil {
		panic(1)
	}

	r.Run(":8080")
}
