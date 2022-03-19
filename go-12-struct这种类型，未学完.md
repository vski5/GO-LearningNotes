# go-12-struct这种类型，又叫结构体

学习自：[Go语言基础之结构体struct - RandySun - 博客园 (cnblogs.com)](https://www.cnblogs.com/randysun/p/15417504.html)

##  struct即结构体，亦被直接称为“结构”。用相关的 不同类型的数据来描述同一个数据对象。

结构体struct充当的就是类的作用，但没有类的继承等概念。

GO用 struct的内嵌 配合 interface 以达到比面向对象更高的拓展性和灵活性。

### struct的定义

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    ......
}
```

类型名 and 字段名 ：

1. 首字母大写则表示公开可被其他包导入就是，外部包可以使用。
2. 首字母小写则表示私有仅在当前的结构体的包内可访问。

同类型的字段可写在同一行。

### 结构体实体化（struct实体化）

结构体实体化才会分配内存，也就是说，实例化才能使用结构体的字段。

struct也是一种类型，所以可以用`var`声明。

```go
var struct实例 struct类型
```



没有**初始化的struct**，其成员变量都是**对应其类型的空（零）值。**

我们通过`.`来访问结构体的字段（成员变量）例如`p1.name`和`p1.age`等,也可以赋值。

```GO
//大写表示外部包可以使用
type person struct {
    name string
    city string
    age  int8
}

func main() {
    var p4 person
    fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}
    
    // 赋值
    // 实例化方式一 键值方式
    //给下面几个子字段重新赋值
    var p1 person
    p1.name = "randy"  //此处我们通过`.`来访问结构体的字段（成员变量）
    p1.city = "ah"
    p1.age = 18
    fmt.Printf("p1=%#v\n", p1) //p4=main.person{name:"randy", city:"ah", age:18} 把p1这三个子字段全打印出来了。
}

```

输出：

```go
DAP server listening at: 127.0.0.1:4306
p4=main.person{name:"", city:"", age:0}
p1=main.person{name:"randy", city:"ah", age:18}

```

### 使用键值对进行初始化

键值对，一对儿，键 对应 结构体的字段，值 对应 该字段的初始值。

例如——>字段name和初始值randy 相当于 键name 值randy

```go
var p1 person
    p1.name = "randy"  //此处我们通过`.`来访问结构体的字段（成员变量）
    p1.city = "ah"
    p1.age = 18
```

等效于：

```go
p1 := person{
    name: "randy",
    city: "ah",
    age:  18,
}
```

————分割线——————

也可以对结构体**指针**进行键值对**初始化**

**初始化**<u>需要初始化所有的字段，填充顺序也要一样。这是对结构体的初始化，不能跟键值初始化混合使用。</u>

```go
p6 := &person{ //这里加&就表示是指针，指向一个储存数值的内存地址。
    name: "ransy",
    city: "ah",
    age:  18,
}
```

等效于：

此处按顺序赋值

```go
p8 := &person{
    "randy",
    "ah",
    28,
}
```

当某些字段没有初始值的时候，该字段可以不写。

此时，没有指定初始值的字段的值就是该字段类型的零（空）值。

### 匿名结构体。

```go
package main
import (
    "fmt"
)
func main() { 
a := struct { 
    name string
    age int
}{name:"mike",age:19} //这里的括号还得在同一行，此处为子字段赋值
fmt.Printf("%#v\n", a)
}
```

### 指针类型结构体

以后再学吧。。。



# go by example的例子

Go 的*结构体(struct)* 是带类型的字段(fields)集合。 这在组织数据时非常有用。

```go
package main

import "fmt"

//这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
	name string
	age  int
}

//newPerson 使用给定的名字构造一个新的 就像是person一样的结构体.以person为模板一样。
func newPerson(nameinput string) *person {
	//符号 * 可以放在一个指针前，如 *intP (此处intp为一个已经声明了的指针)，那么它将得到这个指针指向地址上所存储的值
	//*person指向person的数值。
	//输入值name，返回值*person
	//您可以安全地返回指向局部变量的指针， 因为局部变量将在函数的作用域中继续存在。
	p := person{name: nameinput}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 20}) //使用这个语法创建新的结构体元素。

	fmt.Println(person{name: "Alice", age: 30}) //你可以在初始化一个结构体元素时指定字段名字。

	fmt.Println(person{name: "Fred"}) //省略的字段将被初始化为零值。此处age字段返回0

	fmt.Println(&person{name: "Ann", age: 40}) //& 前缀生成一个结构体指针。

	fmt.Println(newPerson("Jon")) //在构造函数中封装创建新的结构实例是一种习惯用法，按上面封装的函数，输入jon输出&p，也就是&person{name: nameinput，age = 42}

	s := person{name: "Sean", age: 50} //使用 在下面的点 . 来访问结构体字段。
	fmt.Println(s.name)

	sp := &s //也可以对结构体指针使用在下面的点 . 指针会被自动解引用，就是指针本来指向地址现在直接无效化。
	fmt.Println(sp.age)

	sp.age = 51 //结构体是可变(mutable)的。用单个等于号修改结构体内字段名的值
	fmt.Println(sp.age)
}

```

