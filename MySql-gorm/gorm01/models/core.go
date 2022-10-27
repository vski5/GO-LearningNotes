package models

//models是用于操作数据库的包
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	//连接数据库："user:password@tcp(127.0.0.1:3306)/数据库的名字?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:password@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//改gorm.Config这个struct进行设置
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}
	if DB == nil {
		panic("DB IS ERROR")
	}

}

// 赋值给全局的DB共享一个底层连接
var DB *gorm.DB
