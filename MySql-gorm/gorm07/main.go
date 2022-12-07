package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type nameurl struct {
	Username string
	Url      string
}
type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   int
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Access   []Access `gorm:"many2many:role_access"` //
}

func (Manager) TableName() string {
	return "manager"
}

// 一个一对多的表
// 用于表示access和Role之间的连接
type RoleAccess struct {
	AccessId  int
	ManagerId int //RoleId
}

func (RoleAccess) TableName() string {
	return "role_access"
}

type Access struct {
	Id          int
	ModuleName  string //模块名称
	ActionName  string //操作名称
	Type        int    //节点类型 :  1、表示模块    2、表示菜单     3、操作
	Url         string //路由跳转地址
	ModuleId    int    //此module_id和当前模型的id关联       module_id= 0 表示模块
	Sort        int
	Description string
	Status      int
	AddTime     int
	Manager     []Manager `gorm:"many2many:role_access"`
}

func (Access) TableName() string {
	return "access"
}

func main() {
	dsn := "root:密码@tcp(127.0.0.1:3306)/数据库名?charset=utf8mb4&parseTime=True&loc=Local"

	db, err2 := gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true, //打印sql
		//SkipDefaultTransaction: true, //禁用事务
	})

	// DB.Debug()
	if err2 != nil {
		fmt.Println(err2)
	}
	/* 	nameurls := []nameurl{} */
	nameurls222 := []nameurl{}
	/* 	db.Table("manager", "access").Select("manager.username, access.url").
	Joins("left join role_access on manager.role_id = role_access.manager_id").
	Joins("left join access on access.id = role_access.access_id").
	Scan(&nameurls) */
	db.Raw("SELECT manager.username, access.url FROM manager left join role_access on manager.role_id = role_access.manager_id left join access on access.id = role_access.access_id").Scan(&nameurls222)
	/* fmt.Println(nameurls) */
	fmt.Println(nameurls222)
}
