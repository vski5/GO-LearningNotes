package controllers

import "github.com/gin-gonic/gin"

// 核心就是通过实例化结构体（或者用gorm中的where函数），使gorm认识是哪一行。

type UserController struct{}

// 增加数据库行，方法1
func (a UserController) Add(c *gin.Context) {

	c.String(200, "ADD")
}

// 查找数据库行，方法1
func (a UserController) Search(c *gin.Context) {
	c.String(200, "Search")
}

// 删除数据库行，方法1
func (a UserController) Delete(c *gin.Context) {
	c.String(200, "Search")
}
