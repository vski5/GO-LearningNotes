package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 中使用事务，修改&gorm.Config{}
//初始化时禁用 GORM事务 ，这将获得大约 30%+ 性能提升
/*
gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	跳过默认设置
*/

// 本次会用到的结构体
type Bank struct {
	Id      int
	Balance int
}

func (Bank) TableName() string {
	return "bank"
}

func main() {
	r := gin.Default()

	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 改gorm.Config这个struct进行设置
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//SkipDefaultTransaction: true,//初始化时禁用 GORM事务
	})
	if DB == nil {
		panic(1)
	}

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		fmt.Println(err)

	}
	// Id: 1账户减去 100
	u1 := Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
	}
	// panic("遇到了错误")
	// Id: 2账户增加 100
	u2 := Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	// panic("失败")
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()

	}
	tx.Commit()
	r.Run(":8080")
}
