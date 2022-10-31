package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 多表关联查询 多对多

//需要用中间表，在需要关联到的加上，`gorm:"many2many:lesson_student;"`

/*
预加载SQL，在Preload中可以加上匿名函数来return细致的筛选规则。
DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
			return DB.Where("id>2").Order("student.id DESC")
		}).Find(&lessonList)

*/

//DB.Preload("附属的表的名字").Limit(2).Find(&主表写入的地址)

type Lesson struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Student []Student `gorm:"many2many:lesson_student;"`
}

func (Lesson) TableName() string {
	return "lesson"
}

type Student struct {
	Id       int
	Number   string
	Password string
	ClassId  int
	Name     string
	Lesson   []Lesson `gorm:"many2many:lesson_student;"`
}

func (Student) TableName() string {
	return "student"
}

type LessonStudent struct {
	LessonId  int `json:"lesson_id"`
	StudentId int `json:"student_id"`
}

func (LessonStudent) TableName() string {
	return "lesson_student"
}

func main() {
	r := gin.Default()

	dsn := "root:pass@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 改gorm.Config这个struct进行设置
	DB, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if DB == nil {
		panic(1)
	}

	//1、获取所有的学生信息

	r.GET("/preload1", func(c *gin.Context) {
		studentList := []Student{}
		DB.Find(&studentList)
		c.JSON(200, gin.H{
			"result": studentList,
		})
	})

	//2、获取所有的课程信息

	r.GET("/preload2", func(c *gin.Context) {
		lessonList := []Lesson{}
		DB.Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	//3、查询学生信息的时候 展示学生选修的课程

	r.GET("/preload3", func(c *gin.Context) {
		studentList := []Student{}
		DB.Preload("Lesson").Find(&studentList)
		c.JSON(200, gin.H{
			"result": studentList,
		})
	})

	// 4、查询张三 以及张三选修了哪些课程

	r.GET("/preload4", func(c *gin.Context) {

		studentList := []Student{}
		DB.Preload("Lesson").Where("name = ?", "张三").Find(&studentList)
		c.JSON(200, gin.H{
			"result": studentList,
		})
	})

	// 5、查询课程被哪些学生选修了

	r.GET("/preload5", func(c *gin.Context) {
		lessonList := []Lesson{}
		DB.Preload("Student").Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	// 6、查询计算机网络被哪些学生选修了

	r.GET("/preload6", func(c *gin.Context) {
		lessonList := []Lesson{}
		DB.Preload("Student").Where("name=?", "计算机网络").Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	// 7、查询数据指定条件

	r.GET("/preload7", func(c *gin.Context) {
		lessonList := []Lesson{}
		DB.Preload("Student").Offset(1).Limit(2).Order("id desc").Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})

	})
	//
	// 7、查询课程被哪些学生选修的时候去掉id 为1

	r.GET("/preload72", func(c *gin.Context) {

		lessonList := []Lesson{}
		DB.Preload("Student", "id != ?", 1).Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	// 8、 查询课程被哪些学生选修的时候去掉 id 为1或者2的

	r.GET("/preload8", func(c *gin.Context) {

		lessonList := []Lesson{}
		DB.Preload("Student", "id not in (1,2)").Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	// 9 查看课程被哪些学生选修 要求：学生 id 倒叙输出  自定义预加载 SQL

	r.GET("/preload9", func(c *gin.Context) {
		lessonList := []Lesson{}
		DB.Preload("Student", func(db *gorm.DB) *gorm.DB {
			return DB.Where("id>2").Order("student.id DESC")
		}).Find(&lessonList)
		c.JSON(200, gin.H{
			"result": lessonList,
		})
	})

	r.Run(":8080")
}
