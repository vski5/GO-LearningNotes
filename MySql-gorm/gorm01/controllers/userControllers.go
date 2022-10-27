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
func (a UserController) Search(c *gin.Context) {
	/* 	//先实例化,把models.User{}先变成结构体对象的切片，方便models.DB.First(&user)等查询方法赋值
	   	userArry1 := []models.User{}

	   	// 获取所有记录（主键升序）(在实例化的时候，不加匹配条件的情况下)(查询到的数据赋值给user)
	   	models.DB.First(&userArry1)

	   	//测试一下
	   	c.JSON(200, gin.H{
	   		"userArry": userArry1,
	   	}) */
	test := &models.User{}
	models.DB.Find(test)
	c.JSON(200, gin.H{
		"userArry": test,
	})

}

// 删除数据库行，方法1
func (a UserController) Delete(c *gin.Context) {
	c.String(200, "Delete")
}
