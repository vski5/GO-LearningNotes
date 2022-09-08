package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//更细致的用Gin渲染的HTML的讲解，包括传参的方法

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("different/**/*")

	//前台
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "frontend/index.html", gin.H{
			"title": "首页",
			"msg":   " 我是msg",
			"score": 89,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "新闻详情111",
				},
				&Article{
					Title:   "新闻标题222",
					Content: "新闻详情222",
				},
			},
			"testSlice": []string{},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
		})
	})
	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻详情",
		}
		c.HTML(http.StatusOK, "frontend/news.html", gin.H{
			"title": "新闻页面",
			"news":  news,
		})
	})

	//后台
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台首页",
		})
	})
	r.GET("/admin/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin/news.html", gin.H{
			"title": "新闻页面",
		})
	})

	r.Run(":8080")
}
