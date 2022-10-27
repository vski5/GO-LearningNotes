// 用于操作user这个表
package models

// 对应 默认表名是 users ，一般而言命名规则是用表名的单数。
// 这意味着，在连接数据库的时候，默认连接使用名为users的表。
type User struct {
	Id       int //都是表里的列，首字母大写。
	Username string
	Age      int
	Email    string
	AddTime  int
}

// 在结构体上绑定方法 ， 改变结构体的默认表名称
func (U User) TableName() string {
	return "user"
}
