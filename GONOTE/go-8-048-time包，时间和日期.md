# go-8-048-time包，时间和日期

`time`包里有个作为值使用的数据类型`time.Time` ，这个值可以作为 显示和测量时间和日期 的 功能函数。

 **`time.Now()` 获取当前时间，在获得时间的一部分之前要先用这个获取当前时间**

使用 `t.Day()`、`t.Minute()` 、`t.Year()`等等来获取时间的一部分，日月年。

你甚至可以自定义时间格式化字符串，例如： `fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())` 将会输出 `21.07.2011`。

此处的%02d表示只输出两位数，%4d为输出四位数。

```go
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year()) //运行这个首先要获取当下的时间，也就是获得当下时间的前提是有总体的时间。
}
```



Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。

声明变量类型的时候这样——>`var week time.Duration`

Location 类型映射某个时区的时间，

UTC 表示通用协调世界时间。 `t = time.Now().UTC()`



包中的一个预定义函数 `func (t Time) Format(layout string) string` 可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，**看下面的例子理解**

有一些可使用的 预定义的格式，如：`time.ANSIC` 或 `time.RFC822`。 

一般的格式化设计是通过对于一个标准时间的格式化描述来展现的

此处相当于给系统举了个例子，让电脑照着我打出来的例子的格式往里面填数据。

```go
fmt.Println(t.Format("02 Jan 2006 15:04")) 
```

输出：

	21 Jul 2011 10:31





一个例子，输出的结果已经写在每行 `//` 的后面。

示例 4.20 [time.go](examples/chapter_4/time.go)

```go
package main
import (
	"fmt"
	"time"
)

var week time.Duration //先要声明这个变量是一个time.变量里的一个time.Duration类型，Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64。
func main() {
	t := time.Now() //获取当前时间
	fmt.Println(t) // e.g. Wed Dec 21 09:52:14 +0100 RST 2011 就是在输出当前时间
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	// 21.12.2011，按日月年顺数输出
	t = time.Now().UTC() //将t的类型声明，此处就相当于改变了这个time.Now()更加细化了，UTC 表示通用协调世界时间。
	fmt.Println(t) // Wed Dec 21 08:52:14 +0000 UTC 2011
	fmt.Println(time.Now()) // Wed Dec 21 09:52:14 +0100 RST 2011 就是在输出当前时间
	// calculating（计算） times:
	week = 60 * 60 * 24 * 7 * 1e9 // must be in nanosec （转换）
	week_from_now := t.Add(time.Duration(week))
	fmt.Println(week_from_now) // Wed Dec 28 08:52:14 +0000 UTC 2011
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC  
    //包中的一个预定义函数func (t Time) Format(layout string) string可以根据一个格式化字符串来将一个时间 t 转换为相应格式的字符串，你可以使用一些预定义的格式，如time.ANSIC或time.RFC822。
    //func (t Time) Format(layout string) string
	//func指要在func这一部分的代码里面 (t Time此处为time格式的函数的名字) Format(layout string此处填要转变为的格式) string
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	// The time must be 2006-01-02 15:04:05
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
}
```

输出的结果已经写在每行 `//` 的后面。

如果你需要在应用程序在经过一定时间或周期执行某项任务（事件处理的特例），则可以使用 `time.After` 或者 `time.Ticker`：我们将会在第 14.5 节讨论这些有趣的事情。

 另外，`time.Sleep（d Duration）` 可以实现对某个进程（实质上是 goroutine）时长为 d 的暂停。