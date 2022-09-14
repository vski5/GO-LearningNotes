package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

//设置自定义函数---例子1
func UnixToTimeFunc(timestap int) string {
	t := time.Unix(int64(timestap), 0)
	return t.Format("2000-12-15 15:45:56")
}

//设置自定义函数---例子1
func SplicingFunc(input1, input2 string) string {
	return input1 + "-----" + input2 //类java的+字符拼凑
}

func main() {
	r := gin.Default()
	//设置自定义模板函数，要放在初始化后面，html读取的前面
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTimeFunc, //后面是作用的函数自己定义，前面才是用在HTML文件里的
		"Splicing":   SplicingFunc,
	})
	r.LoadHTMLGlob("different/**/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/index.html", gin.H{
			"title": "标题",
			"real": &Article{
				Title:   "test1",
				Content: "test2",
			},
			"date": 1663133960, //时间戳，自定义模板函数的演示
		})
	})

	r.Run(":8080")
}
