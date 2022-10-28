package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// gorm【执行】原生SQL，exec()增删改，Raw()查询再用Scan()写入准备好的golang的空间里（比如切片）
// gorm的SQL生成器

type Ceshi struct { //默认加s，所以对应表ceshis
	Id      int    `json:"id"` //在结构体变为json返回的时候，自动用id替换Id
	Punk    int    `json:"punk"`
	Bigname string `json:"bigname"`
}

var DB *gorm.DB
var err error

func main() {
	//连接数据库
	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	if DB == nil {
		fmt.Println("BD IS error")
	}

	//Raw()查询再用Scan()写入准备好的golang的空间里（比如切片）
	var ceshi111 Ceshi
	DB.Raw("SELECT id, punk, bigname FROM users WHERE id = ?", 3).Scan(&ceshi111)

	//exec()增删改，【执行】原生SQL
	DB.Exec("DROP TABLE users")
	DB.Exec("UPDATE orders SET punk = ? WHERE id IN ?", time.Now(), []int64{1, 2, 3})

}
