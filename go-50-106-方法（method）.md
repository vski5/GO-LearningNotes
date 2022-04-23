# go-50-106-方法（method）


# Struct 就是一些数据放在一起而已，只不过你要给它定一个名字，这个名字就是它的类型签名。


# 方法有什么特别吗？也只是带 receiver(接收器) 的函数罢了吧。

# 接口不就是对一些方法的约束集合吗

# go-50-106-方法（method）
`结构体`很像`类`的简化。

Go `方法`是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。

接受者类型几乎可以是任何类型，因为，任何类型甚至是函数类型都可以有方法，但是接收者**不能是 接口类型 或 指针类型** ，但可以是其他允许类型的指针。

一个类型加上它的方法等价于面向对象中的一个类。

具象化一下：类型 + 方法 = 类

在 Go 中，类型的代码和绑定在它上面的方法的代码不一定要放置在一起，但必须在同一个包里面，举个例子：某些.go文件的都是开篇`package onepack`  。

类型T（或 \*T）上的所有方法的集合叫类型T（或 \*T）的 方法集(method set ) 。

对于一个类型只能有一个给定名称的方法，
但是如果基于接收者类型，是有重载的，具有同样名字的方法可以在 2 个或多个不同的接收者类型上存在，同一个名字的方法作用于不同的接收者类型。
例如：
```go
func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix
```

别名类型没有原始类型上已经定义过的方法。
一个复习：**类型别名 和 类型定义**
```
类型别名 和 原类型 完全一样 ， 但叫起来不一样。
比如说rune和int32是一个类型，但叫起来不一样

底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持，就像是定义 摄氏度 和 华氏度 的两个不同的温度包，底层都是float64但是不可以混在一起表达，
```

**类** 是创建对象的蓝图，描述了所创建的对象**共同**的特性和方法，就像是string这个类里的都是字符串。
类的更严格的定义是由某种特定的元数据所组成的内聚的包，描述了一些对象的行为规则，而这些对象就被称为该类的实例。 

**类有接口和结构。**

**golang里面没有 *类* ，只有struct(结构体)。
struct是没有接口的类

在Go语言中，**类型**分为以下四类： 
- 基本类型：数字，字符串和布尔值属于此类别。 
- 聚合类型：数组和结构属于此类别。 
- 引用类型：指针，切片，map集合，函数和Channel属于此类别。


**定义 *方法* 的一般格式如下**：
```go
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
```
翻译一下就是：
```go
func (接收的名字 接收器的类型) 方法名字(输入的参数) (返回值) {
	...
}
```

在方法名之前，`func` 关键字之后的括号中指定 receiver。

如果 `recv` 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 `object.name` 选择器符号：**recv.Method1()**。

如果 `recv` 是一个指针，Go 会自动解引用。

如果方法不需要使用 `recv` 的值，可以用 **_** 替换它，比如：

```go
func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
```

`recv` 就像是面向对象语言中的 `this` 或 `self`，但是 Go 中并没有这两个关键字。随个人喜好，你可以使用 `this` 或 `self` 作为 receiver 的名字。下面是一个结构体上的简单方法的例子：

示例 10.10 method .go：

```go
package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
```

输出：

    The sum is: 22
    Add them to the param: 42
    The sum is: 7

下面是非结构体类型上方法的例子：

示例 10.11 method2.go：

```go
package main

import "fmt"

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func main() {
	fmt.Println(IntVector{1, 2, 3}.Sum()) // 输出是6
}
```

# 《go in action》上的解答

方法 能给 用户定义的类型 添加 新的行为。
方法实际上也是函数，只是在声明时，在关键字 func 和方法名之间增加了一个参数。
关键字 func 和函数名之间的参数被称作接收者，将函数与**接收者**的**类型**绑在一起。
如果一个函数有接收者，这个函数就被称为方法。
方法有什么特别吗？也只是带 receiver(接收器) 的函数罢了吧。


```go
// 这个示例程序展示如何声明
// 并使用方法
package main
import (

 "fmt"

)

// user 在程序里定义一个用户类型

type user struct {

 name string
 email string

}

  

// notify 使用值接收者实现了一个方法

func (u user) notify() { //notify 方法的接收者被声明为 user 类型的值。需要一个alias（别名）此处为u来获取到这个值。
//这一步还在定义下面的代码块才开始使用。
 fmt.Printf("Sending User Email To %s<%s>\n",

 u.name,

 u.email)

}

  

// changeEmail 使用指针接收者实现了一个方法

func (u *user) changeEmail(endemail string) {

 u.email = endemail

}


// main 是应用程序的入口

func main() {

 // user 类型的值可以用来调用

 // 使用值接收者声明的方法

 bill := user{"Bill", "bill@email.com"} //声明user类型的变量bill

 bill.notify() //此处的bill就相当于上面重新声明出来的的u，充当receiver的作用。使用变量 bill 来 调用 notify 方法

 // 指向 user 类型值的指针也可以用来调用

 // 使用值接收者声明的方法

 lisa := &user{"Lisa", "lisa@email.com"} //这里声明的是user这个类型的指针，依旧可以被作用与user这个类型的方法操作。

 lisa.notify()

 /*

 Go 编译器为了支持这种方法调用背后做的事情: (*lisa).notify()

 把 指针 去解引用为 值， 这样就符合了值接收者的要求。

 */

  

 // user 类型的值可以用来调用

  

 // 使用指针接收者声明的方法

  

 bill.changeEmail("bill@newdomain.com")

 bill.notify()


 // 指向 user 类型值的指针可以用来调用

 // 使用指针接收者声明的方法

 lisa.changeEmail("lisa@newdomain.com")

 lisa.notify()

}
```
虽然没有`类`这个概念的存在，但可以`类实例`只能用`类方法`操作，此处声明的`struct`只能由声明的方法操作。

Go 语言里有两种类型的接收者：值接收者和指针接收者。
- 值接收者：(u user) ，用于调用，
	调用时会使用这个值的一个副本来执行。
- 指针接收者：(u \*user ) ， 用于操作、修改值，
	操作的是一个副本，只不过这次操作的是 从指针指向的值的副本。


## 类型的本质

- 如果是要创建一个新值：该类型的方法就使用值接收者。
- 如果是要修改当前值：就使用指针接收者。
	就像是上文中的：
	
	   func (u *user) changeEmail(endemail string) {
			u.email = endemail
	    }


这个答案也会影响程序内部传递这个类型的值的方式：是按值做传递，还是按指针做传递。
保持传递的一致性很重要。
这个背后的原则是，不要只关注某个方法是如何处理这个值，而是要关注 这个值的本质是什么
