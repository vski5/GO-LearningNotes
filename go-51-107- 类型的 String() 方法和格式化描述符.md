# go-51-107- 类型的 String() 方法和格式化描述符
使用 `String()` 方法来定制类型的字符串形式的输出，换句话说：一种可阅读性和打印性的输出。

如果类型定义了 `String()` 方法，它会被用在 `fmt.Printf()` 中生成默认的输出：等同于使用格式化描述符 `%v` 产生的输出。

`fmt.Print()` 和 `fmt.Println()` 也会自动使用 `String()` 方法。

示例 10.22 method_string.go：

```go
package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
```

输出：

    two1 is: (12/10)
    two1 is: (12/10)
    two1 is: *main.TwoInts
    two1 is: &main.TwoInts{a:12, b:10}

这个代码删点东西也是行得通的：
```go
package main
import (

 "fmt"

)
type TwoInts struct {

 a int

 b int

}
func main() {

 two1 := new(TwoInts)

 two1.a = 12

 two1.b = 10

 fmt.Printf("two1 is: %v\n", two1)

 fmt.Println("two1 is:", two1)

 fmt.Printf("two1 is: %T\n", two1) //格式化描述符 %T 会给出类型的完全规格

 fmt.Printf("two1 is: %#v\n", two1) //%#v 会给出实例的完整输出，包括它的字段
}
```


**练习 10.12** type_string.go

给定结构体类型 T:

```go
type T struct {
    a int
    b float32
    c string
}
```

值 `t`: `t := &T{7, -2.35, "abc\tdef"}`。给 T 定义 `String()`，使得 `fmt.Printf("%v\n", t)` 输出：`7 / -2.350000 / "abc\tdef"`。
```go

package main

  

import (

 "fmt"

)

  

type T struct {

 a int

 b float32

 c string

}

  

func main() {

 t := &T{7, -2.35, "abc\\tdef"}

  

 fmt.Printf("%v\n", *t)

}
```



练习 10.13 celsius.go

为 float64 定义一个别名类型 Celsius，并给它定义 String()，它输出一个十进制数和 °C 表示的温度值。

```go

package main

  

import (

 "fmt"

)

  

type Celsius float64

type float64 struct {

 shijingzhi float32

 wendu float32

}

  

func (ce Celsius) String() {

 fmt.Printf(" 十进制：%v \n 温度 %v °C \n",

 ce.shijingzhi,

 ce.wendu)

}

  

func main() {

 text1 := Celsius{12, 26}

 text1.String()

}
```


**练习 10.14** days.go

为 int 定义一个别名类型 `Day`，定义一个字符串数组它包含一周七天的名字，为类型 `Day` 定义 `String()` 方法，它输出星期几的名字。使用 `iota` 定义一个枚举常量用于表示一周的中每天（MO、TU...）。
```go

package main

  

import (

 "fmt"

)

  

type Day int

  

const (

 Monday Day = iota

 Tuesday

 Wednesday

  

)

  

func (d Day) String() string {

 return [...]string{"Monday", "Tuesday", "Wednesday"}[d]

}

  

var d Day = Tuesday

  

func main() {

 fmt.Print(d)

 switch d {

 case Monday:

 fmt.Println(" goes up.")

 case Tuesday:

 fmt.Println(" goes down.")

 default:

 fmt.Println(" stays put.")

 }

}
```

