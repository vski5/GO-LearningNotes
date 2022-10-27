package controllers

import (
	"gorm01/models"

	"github.com/gin-gonic/gin"
)

// 核心就是通过实例化结构体（或者用gorm中的where函数），使gorm认识是哪一行。

type UserController struct{}

// 增加数据库行，方法1
func (a UserController) Add(c *gin.Context) {

	//先实例化
	//user := &User{}
	//，把获取的数据放结构体里面。
	//models.DB.Find(&user)
	c.String(200, "add")
}

// 查找数据库行，方法1 ，默认查找第一行
/* find和first用同一套筛选方法，但是范围不同，first只查询第一条，find都查询 */
func (a UserController) Search(c *gin.Context) {
	//先实例化,把models.User{}先变成结构体对象的切片，方便models.DB.First(&user)等查询方法赋值
	userArry1 := []models.User{}
	userArry2 := []models.User{}
	userArry3 := []models.User{}
	userArry4 := []models.User{}

	// 获取第一行记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
	models.DB.First(&userArry1)

	//主键是数字类型，可以使用 内联条件 来检索对象
	models.DB.First(&userArry2, 2)
	// SELECT * FROM users WHERE id = 2;

	models.DB.First(&userArry3, "2")
	// SELECT * FROM users WHERE id = 2;

	models.DB.Find(&userArry4, []int{1, 2, 3})
	// SELECT * FROM users WHERE id IN (1,2,3);

	//测试一下
	c.JSON(200, gin.H{
		"userArry1": userArry1,
		"userArry2": userArry2,
		"userArry3": userArry3,
		"userArry4": userArry4,
	})
}

// 查找数据库行，方法2 ，默认查找 所有数据
func (a UserController) Search2(c *gin.Context) {
	//先实例化,把models.User{}先变成结构体对象的切片，方便models.DB.First(&user)等查询方法赋值
	userArry1 := []models.User{}
	userArry2 := []models.User{}

	// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
	models.DB.Find(&userArry1)

	//获取特定数据
	models.DB.Where("id = ?", "2").Find(&userArry2)

	//测试一下
	c.JSON(200, gin.H{
		"userArry1": userArry1,
		"userArry2": userArry2,
	})
}

// 查找数据库行，方法3 ，指定结构体查询字段
func (a UserController) Search3(c *gin.Context) {
	//先实例化,把models.User{}先变成结构体对象的切片，方便models.DB.First(&user)等查询方法赋值
	userArry1 := []models.User{}
	userArry2 := []models.User{}
	userArry3 := []models.User{}

	// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
	models.DB.Find(&userArry1)

	models.DB.Where(&models.User{Username: "gorm"}, "username", "Age").Find(&userArry2)
	// SELECT * FROM users WHERE Username = "gorm" AND age = 0;

	models.DB.Where(&models.User{Username: "gorm"}, "Age").Find(&userArry3)
	// SELECT * FROM users WHERE age = 0;

	//获取特定数据
	models.DB.Where("id = ?", 2)

	//测试一下
	c.JSON(200, gin.H{
		"userArry1": userArry1,
		"userArry2": userArry2,
		"userArry3": userArry3,
	})
}

// 删除数据库行，方法1
func (a UserController) Delete(c *gin.Context) {
	c.String(200, "Delete")
}
