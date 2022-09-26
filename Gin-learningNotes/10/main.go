package main

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取、解析XML数据，并且绑到结构体上去.用到c.GetRawData()获取c.Request.body数据，xml.Unmarshal(xmlByte, &article)反序列化xmlByte,并且绑定到Article结构体

//用Postman去post下面的xml数据
/*
<?xml version="1.0" encoding="UTF-8"?>
<article>
<content type="string">我是内容</content>
<title type="string">我是标题</title>
</article>
*/

//结构体：
type Article struct {
	Content string `xml:"content" json:"content"`
	Title   string `xml:"title" json:"title"`
}

func main() {
	r := gin.Default()
	r.POST("/XML", func(c *gin.Context) {
		//1.实例化结构体
		article := &Article{}
		//2.使用c.GetRawData(),获得c.Request.Body中的内容，属于是XML到byte[]了
		xmlByte, _ := c.GetRawData()
		//打印xmlByte，看看长什么样
		fmt.Println(xmlByte)
		//3.反序列化xmlByte,并且绑定到Article结构体
		err := xml.Unmarshal(xmlByte, &article)
		//4.判断反序列化是否成功
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(200, article)
		}

	})
	r.Run(":8080")
}
