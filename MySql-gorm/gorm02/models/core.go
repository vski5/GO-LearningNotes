package models

//models是用于操作数据库的包
import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	//连接数据库："user:password@tcp(127.0.0.1:3306)/数据库的名字?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	//改gorm.Config这个struct进行设置
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
