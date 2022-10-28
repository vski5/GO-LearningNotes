package controllers

import (
	"gorm01/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 核心就是通过实例化结构体（或者用gorm中的where函数），使gorm认识是哪一行。

type UserController struct{}

// 增加数据库行，方法1
func (a UserController) Add(c *gin.Context) {
	timeUnix := time.Now().Unix()
	timeUnixInt64 := strconv.FormatInt(timeUnix, 10)
	timeUnixInt, _ := strconv.Atoi(timeUnixInt64)
	user := models.User{
		Id:       77,
		Username: "ccsssaaa",
		Age:      33,
		Email:    "rvr@ewc1ew1c",
		AddTime:  timeUnixInt,
	}

	models.DB.Create(&user) // 通过数据的指针来创建

	//user.Id             // 返回插入数据的主键
	//result.Error        // 返回 error
	//result.RowsAffected // 返回插入记录的条数
	c.String(200, "add")
}

// 增加数据库行，方法2 ，根据map创建
// 根据 map[string]interface{} 和 []map[string]interface{}{} 创建记录
func (a UserController) Add2(c *gin.Context) {
	timeUnix := time.Now().Unix()
	timeUnixInt64 := strconv.FormatInt(timeUnix, 10)
	timeUnixInt, _ := strconv.Atoi(timeUnixInt64)

	models.DB.Model(&models.User{}).Create(map[string]interface{}{
		"Id":       77,
		"Username": "ccsssaaa",
		"Age":      33,
		"Email":    "rvr@ewc1ew1c",
		"AddTime":  timeUnixInt,
	})

	// batch insert from `[]map[string]interface{}{}`
	models.DB.Model(&models.User{}).Create([]map[string]interface{}{
		{
			"Id":       23,
			"Username": "ccsssaaa",
			"Age":      33,
			"Email":    "rvr@ee1c",
			"AddTime":  timeUnixInt,
		},
		{
			"Id":       72,
			"Username": "ccsssaaa",
			"Age":      33,
			"Email":    "rvr@ewcw1c",
			"AddTime":  timeUnixInt,
		},
	})

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

	// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
	models.DB.Find(&userArry1)

	//用结构体查特定数据
	userArry2 := &models.User{Id: 3}
	models.DB.Find(userArry2)

	//获取特定数据
	//models.DB.Where("id = ?", 2)  ,  符号很多种
	//models.DB.Where("id > ?", 2)
	userArry3 := []models.User{}
	models.DB.Where("id in (?)", []int{1, 2, 3}).Find(&userArry3)
	//models.DB.Find(&userArry4, []int{1, 2, 3}) ，也可以在find后面写

	//模糊查询，查询包含关键字的
	userArry4 := []models.User{}
	models.DB.Where("id like ?", "%关键字%").Find(&userArry4)

	//测试一下
	c.JSON(200, gin.H{
		"userArry1": userArry1,
		"userArry2": userArry2,
		"userArry3": userArry3,
	})
}

// 删除数据库行，删除一条记录时，删除对象需要指定主键，否则会触发 批量 Delete
// 方法1
func (a UserController) Delete(c *gin.Context) {
	// 全删了
	//models.DB.Delete(&models.User{})

	// 带额外条件的删除
	models.DB.Where("name = ?", "jinzhu").Delete(&models.User{
		Id: 73, //删除id为73的
	})

	c.String(200, "Delete")
}

// 更新数据库行， 方法1
// Save 会保存所有的字段，即使字段是零值
func (a UserController) Update1(c *gin.Context) {

	user := models.User{Id: 73}
	models.DB.Find(&user)
	user.Username = "gogogogo"
	user.Age = 1
	models.DB.Save(&user)

	c.String(http.StatusOK, "Edit")
}
