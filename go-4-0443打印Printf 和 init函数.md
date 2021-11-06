# 打印Printf 和 init函数

**一个tip 首字母大写的函数可以在包外使用**

Printf函数的首字母就是大写的。

```go
fmt.Print("Hello:", 23)
```

将输出：`Hello: 23`。

这一句可以理解为

```go
fmt.Print("一句话，因为在双引号内所以打印出来的是原文不计算", %某)
//“%某”是指明了格式，Print会把这一部分给格式化了，要格式化成什么格式，就取决于“%某”指明了什么格式。
```

当一个名为a的变量已经被赋值了，那么`a := 20` 是不被允许的，

但是 `a = 20` 是可以的，因为这是给相同的变量赋予一个新的值。

你要是在局部变量（就是一个花括号里的字符串）中声明了这个a，但你不去用，那就会报错。

### **并行** 或 **同时** 赋值

同一类型的多个变量可以声明在同一行，如：

```go
var a, b, c int
```

(这是将类型写在标识符后面的一个重要原因)

多变量可以在同一行进行赋值，

当变量a，b，c没被赋值时，这里用的是`:=`，看清楚，是带引号的。

```go
a, b, c := 5, 7, "abc"
```

假设了变量 a，b 和 c 都已经被声明，只用一个`=`，想到于，就算是有值了，我也要用新值取而代之。

```go
a, b, c = 5, 7, "abc"
```

右边的这些值以相同的顺序赋值给左边的变量，所以 a 的值是 `5`， b 的值是 `7`，c 的值是 `"abc"`。



并行赋值也被用于当一个函数返回多个返回值时，

这里把两个值赋给了一个函数

这里的 `val` 和错误 `err` 是通过调用 `Func1` 函数同时得到：

```
val, err = Func1(var1)
```



```go
var 变量名 类型 = 表达式
```

GO语言很智能，不写类型也能通过表达式认出来（**大概？**）



### 交换两个变量的值

```go
a, b = b, a
```

名为a和名为b的变量就被交换了。

### 如何抛弃值

因为Go语言需要使用所有被声明的变量，但我并不需要得到一个函数的返回值，那么我就可以抛弃这个被赋予的值。

```go
_,b=5,7
```

  下划线让5被抛弃了。

下划线_是一个只写变量，它存在，但不完全存在，只是用来写出来，可以说是占个位子。

# init函数

init函数可以初始化变量。执行方式是init在每个包完成初始化之后自动执行，init比main函数的优先级还高。

main就是那个每个可执行文件有且只有一个的包。

一个例子。

init 函数中计算变量 Pi 的初始值

```go
package trans

import "math"

var Pi float64  //声明名为pi的变量的类型为float64

func init() {
   Pi = 4 * math.Atan(1) // init() function computes Pi
    //4后面的*是乘号的意思，atan( ) 函数，就一求tan值的，括号里的1就是求tan（1）等于多少。
    //说白了就是在求pi=4*tan1
}
```



另外一个例子。

init.go中导入了包 trans

需要init.go目录为./trans/init.go

并且使用到了变量 Pi：

```go
package main

import (
   "fmt"
   "./trans"  //  init.go目录为./trans/init.go ，这里的 ./trans 就是调用这个目录下的包
)

var twoPi = 2 * trans.Pi   //声明函数

func main() {
   fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}          
```



init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine，如下面这个例子当中的 `backend()`：

```go
func init() {    // setup preparations 这里用init（）函数功能就是在程序开始之前直接用init（）函数调用后台中的一个功能goroutine
   go backend()
}
```





printf一定要用占位符，在引号里面。

一个例子

```
func main() {

  var a string = "12345"

  fmt.Printf("string=%v\n", a)

}
```

输出结果是

```
string=12345 
```



要是没有在引号里面的占位符号的话

```
func main() {
	var a Rope = "12345"
	fmt.Printf("string=", a)
}
```

输出结果是

```
string=%!(EXTRA main.string=12345)
```

