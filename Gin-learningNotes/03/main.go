package main

//返回不同类型的值。c.String() c.JSON() c.JSONP() c.XML() c.HTML()
import (
	"github.com/gin-gonic/gin"
)

type Stext struct {
	Key1 string
	Key2 int
}

func main() {
	r := gin.Default()

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
		s := Stext{ //可以用&Stext
			Key1: "value1",
			Key2: 2,
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

	r.Run(":8080")
}
