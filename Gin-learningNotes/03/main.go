package main

//返回不同类型的值。c.String() c.JSON() c.JSONP() c.XML() c.HTML()
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Stext struct {
	Key1 string `json:"key1,GET之后会显示标签内的内容,仅限改成小写字符"`
	KEY2 int    `json:"key2"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("default/*") //初始化HTML文件地址，记得写在最前面
	/*   1.返回string*/
	r.GET("/string", func(c *gin.Context) {
		aid := c.Query("aid")
		c.String(200, "返回值为?aid=%v", aid)
	})

	/*   2. 返回JSON*/
	/*用gin.H{"key":"value"}自己拼接JSON
	gin.H{"key":  "value"}就是map["string"]interface{}{"key":  "value"}
	*/
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"key":  "value",
			"key2": "value2",
		})
	})
	/*用结构体拼装*/
	r.GET("/structjson", func(c *gin.Context) {
		/*用struct写一个JSON,记得首字母大写，不然传不进去值*/
		var structjson struct {
			Key1 string
			Key2 int
			Key3 float64
		}
		structjson.Key1 = "value1"
		structjson.Key2 = 666
		structjson.Key3 = 999
		c.JSON(200, structjson)
		/*结构体的定义可以在外面定义，但记得大写以便传值,使用s := &structname 来实例化*/
		s := &Stext{ //可以不用&，但会浪费内存
			Key1: "value1",
			KEY2: 2,
		}
		c.JSON(200, s)

	})

	/*   3. 返回JSONP*/
	/*和JSON差不多，但是格式不同，类似c.query的aid,网站为/JSONP?callback=x ， 此处的x可为其他的值
	 */
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo":  "bar",
			"key2": "value2",
		}
		c.JSONP(200, data)
	})
	// 将输出：x({"foo":"bar","key2":"value2"});

	/*   4. 返回XML*/
	r.GET("/xml", func(c *gin.Context) {
		type Msg2 struct {
			Key  string
			Key2 string
		}
		var ms Msg2
		ms.Key = "value1"
		ms.Key2 = "VALUE2"
		c.XML(200, ms)
	})

	/*HTML的渲染*/
	/*需要先初始化r.LoadHTMLGlob("default/*")确定文件的位置*/
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html",
			gin.H{
				"title": "我是商品页面",
				"price": 20,
			})
	})

	r.Run(":8080")

}
