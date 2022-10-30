package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 多表关联查询 一对一、 一对多
// has many 与另一个模型建立了一对多的连接。 不同于 has one，拥有者可以有零或多个关联模型。

//DB.Preload("附属的表的名字").Limit(2).Find(&主表写入的地址)

// `gorm:"foreignKey:侧键;references:主键"`
// 不写会默认主键是Id，侧键是连上去的struct的名字加id，比如连接Cate这个表（struct）会默认CateId是侧键。最好写一下。
type Article struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Description int         `json:"description"`
	CateId      string      `json:"cate_id"`
	State       int         `json:"state"`
	ArticleCate ArticleCate `gorm:"foreignKey:CateId;references:Id"` //`gorm:"foreignKey:侧键;references:主键"`
}

type ArticleCate struct { //	`gorm:"references:Id"`
	Id    int    `json:"id"`
	Title string `json:"title"`
	State int    `json:"state"`
}

/* 从侧链表ArticleCate查询连接到的主表Article */
type Article222 struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description int    `json:"description"`
	CateId      string `json:"cate_id"`
	State       int    `json:"state"`
}

type ArticleCate222 struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	State   int       `json:"state"`
	Article []Article `gorm:"foreignKey:CateId;references:Id"`
}

func (Article222) TableName() string {
	return "article"
}
func (ArticleCate222) TableName() string {
	return "article_cate"
}
func (Article) TableName() string {
	return "article"
}
func (ArticleCate) TableName() string {
	return "article_cate"
}

func main() {
	r := gin.Default()

	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 改gorm.Config这个struct进行设置
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DB == nil {
		panic(1)
	}
	r.GET("/preload1", func(c *gin.Context) {
		var articleList []Article
		DB.Preload("ArticleCate").Limit(2).Find(&articleList)
		c.JSON(200, gin.H{
			"把侧链表ArticleCate连接到主表Article": articleList,
		})
	})
	r.GET("/preload2", func(c *gin.Context) {
		var articleCate222List []ArticleCate222
		DB.Preload("Article222").Limit(2).Find(&articleCate222List)
		c.JSON(200, gin.H{
			"从侧链表ArticleCate查询连接到的主表Article": articleCate222List,
		})
	})

	r.Run(":8080")
}
