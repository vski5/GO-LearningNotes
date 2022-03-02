# go-14-053-switch 结构

 `var`关键字用来定义变量

​	变量 `var1` 可以是任何类型，而 `val1` 和 `val2` 则可以是同类型的任意值。

​	<u>这里提到的三个变量都是同类型的变量。</u>

*default*是默认的意思，在前面的`case`都不满足`var1`（也就是`switch`后面跟的值）的时候，就输出`default`后面的代码。

每一个 *case* 分支都是唯一的,从上至下逐一测试,直到匹配为止。`case` 与 `case` 之间是独立的代码块。

```go
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}	
```

格式的要求——前花括号 `{` 必须和 switch 关键字在同一行。

一个例子

```go
package main

import (
    "fmt"
    "time"
)

func main() {

    i := 2
    fmt.Print("write ", i, " as ")
    switch i { //switch带了表达式i，也就是名为i的一个变量
    case 1: //逐一测试，直到case后的变量和switch后的变量（表达式）一致为止，然后输出下面的东西
        fmt.Println("one")  //1，2，3这些条件的数据格式要为相同的格式。
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

    switch time.Now().Weekday() { //调用time函数的时候，要先把time.Now()给求出来，然后再搞出来小的部分，比如t.Hour()。在这里time.Now().Weekday()
    case time.Saturday, time.Sunday: //这里两个条件之间有个逗号，只要满足任意一个条件就行。可以有很多个逗号来隔开条件。但这些条件的数据格式要为相同的格式
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch { //这里不带表达式的 switch 是实现 if/else 逻辑的另一种方式
    case t.Hour() < 12: //此处相当于if t.Hour() < 12
        fmt.Println("It's before noon")
    default:  //此处类似于else
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {//interface{}类型是没有方法的接口，空接口。下面有关于这一行的详细解释
        //类型开关 () ,也就是switch结构
        //此处比较类型，也就是核对case后面的类型和switch后面的类型是否一致。可以用来找出来一个接口值的类型。
        //专业的说法把这个操作称为type switcht
        switch t := i.(type) {//i.(type)是求得i的数据类型的方法
        case bool: 
            //在这个例子中，变量 在每个分支中都列出来了相应的类型。
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true) //这里的whatAmI（），相当于一个简写的上面的 whatAmI := func(i interface{}){}，和上面的功能一样。因为whatAmI已经被赋值成为了一种函数。
    whatAmI(1)
    whatAmI("hey")
}
```

最后输出

```go
write 2 as two
It's the weekend
It's after noon
I'm a bool
I'm an int

```



