# go-45-100-结构（struct）

# 结构体 （struct）

Go 通过**类型别名**（alias types）和**结构体**的形式支持用户自定义类型，或者叫定制类型。

一个带属性的结构体试图表示一个现实世界中的实体。

结构体是复合类型（composite types），**结构体是一种自定义的数据类型**

当需要定义一个类型，它由一系列属性组成，

例如：

```go
type person struct {
    name string
    age  int
}
```

每个属性都有自己的类型和值的时候，就应该使用结构体，

它把数据聚集在一起，然后可以访问这些数据，就好像它是一个独立实体的一部分。



Go 的*结构体(struct)* 是带类型的字段(fields)集合。 这在组织数据时非常有用

```go
package main

import "fmt"

//这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

/*
1. 对于指针而言，加上*可知指针指向的地址储存的值是什么。
2. 对于变量而言，加上&就能知道变量在内存的哪儿
*/
//newPerson 使用给定的name构造一个新的 person 结构体.结构体是一种自定义的数据类型
//返回值为*person，人造类型person的指针。
func newPerson(name1 string) *person {
	//您可以安全地返回指向局部变量的指针， 因为局部变量将在函数的作用域中继续存在。
	p := person{name: name1}
	p.age = 42
	return &p //返回person所在的位置。&p就是*person的值
}

func main() {
	//使用这个语法创建新的结构体元素。
	fmt.Println(person{"Bob", 20})
	//你可以在初始化一个结构体元素时指定字段名字。
	fmt.Println(person{name: "Alice", age: 30})
	//省略的字段将被初始化为零值。
	fmt.Println(person{name: "Fred"})
	//& 前缀生成一个结构体指针。
	fmt.Println(&person{name: "Ann", age: 40})
	//在构造函数中封装创建新的结构实例是一种习惯用法
	fmt.Println(newPerson("Jon")) //这一步与上面自己构建的方法 newPerson() 联系起来了，只要是用newPerson()就会将age赋值为p.age = 42
	//使用.来访问结构体字段。
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	//也可以对结构体指针使用. - 指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)
	//结构体是可变(mutable)的。
	sp.age = 51
	fmt.Println(sp.age)
}

```

输出：

```
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```

## new()与结构体(struct)

**一个复习**

make()的使用方法是：

```go
func make([]T, len, cap)
/*
func make([]类型，长度，最大长度)
*/
```

new()的使用方法是：

```go
new([100]int)[0:50]
/*
func new([最大长度]类型)[起始位置:终结位置]
*/
```

**使用 new**

new为新结构体分配内存，返回已经被分配内存的指针：

```go
var t *T = new(T)
```

拆开就是一般用法：

```go
var t *T //T可以说一个被命名的struct
t = new(T)
```

变量`t`是指向`T`的指针。

`t` 通常被称做类型 T 的一个实例（instance）或对象（object）。

##  **选择器（selector）**

**选择器**：就像在面向对象语言所作的那样，可以使用点号符给字段赋值：`structname.fieldname = value`。

同样的，使用点号符可以获取结构体字段的值：`structname.fieldname`。

无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的 **选择器符（selector-notation）** 来引用结构体的字段：

```go
type myStruct struct { i int }
var v myStruct    // v是结构体类型变量
var p *myStruct   // p是指向一个结构体类型变量的指针
v.i
p.i
```

初始化结构体实例，或者说是创造一个结构体实例的简写方法：

```go
ms := &struct1{1,2,"c"}
// &struct1{a, b, c} 是一种简写，底层仍然会调用 `new ()`，这里值的顺序必须按照字段顺序来写。
//或者
var ms struct1
ms = struct1{1,2,"c"}
```

表达式 `new(Type)` 和 `&Type{}` 是等价的。

# 递归结构体 与 链表 与 二叉树

struct可以引用自身来定义，这会用在 **自定义 <u>链表</u>**，或，**<u>二叉树</u> 的元素**（又叫二叉树的结构体）。

struct引用自身来定义的时候，节点包含指向临近节点的连接。

## 链表：

![img](file://C:\Users\%E6%9D%8E%E5%B0%A4%E7%9A%84%E5%85%89%E5%BD%B1%E7%B2%BE%E7%81%B5.LAPTOP-85F2O9KH\Desktop\GOnotes\the-way-to-go_ZH_CN-master\eBook\images\10.1_fig10.3.jpg?raw=true?lastModify=1648974556)

这块的 `data` 字段用于存放有效数据（比如 float64），`su` 指针指向后继节点。

Go 代码：

```go
type Node struct {
    data    float64
    su      *Node
}
```

链表中的第一个元素叫 `head`，它指向第二个元素；

最后一个元素叫 `tail`，它没有后继元素，所以它的 `su` 为 nil 值。

当然真实的链接会有很多数据节点，并且链表可以动态增长或收缩。

同样地可以定义一个双向链表，它有一个前趋节点 `pr` 和一个后继节点 `su`：

```go
type Node struct {
    pr      *Node
    data    float64
    su      *Node
}
```

## 二叉树：

二叉树：

![](C:\Users\李尤的光影精灵.LAPTOP-85F2O9KH\Desktop\GOnotes\the-way-to-go_ZH_CN-master\eBook\images\10.1_fig10.4.jpg)

二叉树中每个节点最多能链接至两个节点：左节点（le）和右节点（ri）.

与双向链表不同，这是向下分层的结构。

这两个节点本身又可以有左右节点，依次类推。

树的顶层节点叫根节点（**root**），底层没有子节点的节点叫叶子节点（**leaves**），叶子节点的 `le` 和 `ri` 指针为 nil 值。在 Go 中可以如下定义二叉树：

```go
type Tree struct {
    le      *Tree
    data    float64
    ri      *Tree
}
```

## 结构体转换

alias：别名

当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型，它们可以如示例 10.3 那样互相转换，

同时需要注意其中非法赋值或转换引起的编译错误。

示例 10.3：

```GO
package main

import "fmt"

type number struct {
	f float32
}

type nr number // alias type

func main() {
	a := number{5.0}
	b := nr{7.0}
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var c number = b    // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	var c = number(b)
	var d = nr(a)
	fmt.Println(a, b, c, d)
}

```

输出：

    {5} {7} {7} {5}



# 

