# go-25-067-将函数作为参数 与 Strings.Indexfunc()

函数作为参数，可在其他函数内调用执行，这种操作被称为**回调**

在函数调用中，使用 类似于：

```go
fucn f(x int , y func(int,int)){
}
```

其中的 `y func(int,int)`就是指输入值`y`的数据类型是函数，且函数`y`的输入值类型为两个`int`

```go
package main
func ("fmt")
func main(){
    callback(1,Add)
}
func Add(a,b int){
    func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

```

输出：

	The sum of 1 and 2 is: 3



在 Go 语言中，在一个字符串中查找满足特定条件字符-->用Strings.Indexfunc()

格式为：

```go
func IndexFunc(s string, f func(rune) bool) int
```

第一个参数`s`为源字符串，

第二个参数`f`是一个匿名函数，为bool型，返回ture则表示此字符为我们要查找出来的函数，用来判断要被检索的字符符不符合条件。

rune 是**Go 中的内置类型**，它是 int32 的别名。 Rune 代表 Go 中的 Unicode 代码点。

**<u>返回值类型为`int`，如果符合条件就返回第一次出现要被检索的字符的索引，不符合条件就返回-1.</u>**

一个例子：

```go
package main

import (
	"fmt"
	"strings"
)

func checkRune(r rune) bool { //rune就是int32，后面的bool指这个函数作用是判断是否。
	if r > 73 {
		return true
	} else {
		return false
	}
}
func main() {
	//使用 Strings.IndexFunc() 函数，查找字符串中满足特定条件的字符
	stringsexample := "I love Golang and I study Golang From hacker"
	indexFunc := strings.IndexFunc(stringsexample, checkRune)
	//stringsexample就是要被查验的字符串
	//checkRune就是上面自造的函数：func checkRune(r rune) bool ，符合条件就返回第一次出现要被检索的字符（也就是stringsexample）的索引.
	//查找的条件是字符的 ASCII 码大于 73 的字符第一次出现的位置，此处的输入值r rune相当于将stringsexample变成了int32类型。
	fmt.Println("indexFunc =", indexFunc)
}

```

另外的例子：

用Strings.IndexFunc()函数，查找字符串中满足条件的字符。

```go
package main

import (
	"fmt"
	"strings"
)

func checkRune(r rune) bool {
	if r == 'G' || r == 'a' { //  ||是逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
		return true
	}
	return false
}
func main() {
	//使用 Strings.IndexFunc() 函数，查找字符串中满足特定条件的字符
	strHaiCoder := "I love Golang and I study Golang From HaiCoder"
	indexFunc := strings.IndexFunc(strHaiCoder, checkRune)
	fmt.Println("indexFunc =", indexFunc)
}

```

输出：

```
indexFunc = -1

```



