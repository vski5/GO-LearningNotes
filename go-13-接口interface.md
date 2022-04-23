# go-13-接口interface

理解下面的代码块的准备工作

1. `interface`不单单是在声明这是接口，也在声明{}里面的东西是方法，里面东西的格式是，`名字() 类型`

2. 一种切片的声明方法：

   ```go
   animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
   ```

   此处的格式为：

   ```GO
   名字 := []接口{类型{}，类型{}，类型{}}
   ```

3. 在这个例子中，类型就是用来被接口里的方法作用的

```GO
package main

import "fmt"

type Animal interface { //声明一个名为Animal的接口，其中包括名为Speak()的方法，此处不单是声明接口，也创造了声明了方法
	Speak() string
}

type Dog struct { //声明Dog是一种struct（类型）
}

func (Dog) Speak() string { //声明：将类型Dog用方法Speak() 作用的时候，返回{}内的东西
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string { //var c Cat，下同
	return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}
func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals { //for : range 是历遍函数，这个_,animal 前一个_舍去用来计数的系数，
		//后面的animal被赋值为animals，上一行已经创建 Animals 切片，并将每种类型中的一种放入该切片中
		fmt.Println(animal.Speak()) //注意这一行还在for:=range函数的{}内，所以说这一行的函数就是，被历遍的函数，被依次代入了上面创造的切片的东西。

	}
	fmt.Println(animals) //这里打印出来的就是[{} {} {} {}]，因为没有被"方法"作用的struct(类型)，返回了一个空集。
	//上面能返回东西，是因为已经赋值过了，看一下包含return的代码块就知道了。

}

```



一个抽象：

```go
type 接口名称 interface { //声明一个接口，其中包括名为的Speak()方法，此处不单是声明接口，也创造了声明了方法。
    //接口中的所有方法都要被实现。
	Speak() string
}

type Dog struct { //声明Dog是一种struct（类型）
}

func (Dog) Speak() string { //声明：将类型Dog用方法Speak() 作用的时候，返回{}内的东西
	return "Woof!"
}

//1-声明接口，其中有什么方法；
//2-声明类型；
//3-封装方法，
//func (类型名称) 方法名称() 类型 { 
//	return "返回值"
//}

```



##  `interface{}` 类型





**方法签名**由  **方法名称**  和  **一个参数列表**  （方法的参数的顺序和类型）组成。 

**就像是一个目录，记录着方法叫什么，在哪儿是什么怎么调用。**

也就是方法的签名,来区分不同方法的标示符。

注意，方法签名不包括方法的返回类型。不包括返回值和访问修饰符。

方法签名的集合叫做：_**接口**(Interfaces)_。

***接口***定义了对象的行为，接口指定类型应具有的方法，类型决定如何实现这些方法。

interface{}类型是没有方法的接口，空接口。

## 一个例子

```GO
package main

import "fmt"

type Animal interface {
	Eat()
	Run()
}

type Dog struct {
	Name string
}

func (d *Dog) Eat() {
	fmt.Printf("%s is eating\n", d.Name)
}

func (d *Dog) Run() {
	fmt.Printf("%s is running\n", d.Name)
}

func ShowEat(animal Animal) {
	animal.Eat()
}

func ShowRun(animal Animal) {
	animal.Run()
}

func main() {
	dog := Dog{Name:"Kenny"}
	ShowEat(&dog)
	ShowRun(&dog)
}
```

输出

```GO
Kenny is eating
Kenny is running
```



**<u>gobyexample里面一个关于接口的例子：</u>**

一些英语：

area——面积

geometry——几何学

width——宽度

```go
package main

import (
	"fmt"
	"math"
)

type geometry interface { //这是一个几何体的基本接口,interface就是接口的意思，此处声明geometry的类型为interface（接口）
	area() float64
	perim() float64 //声明这两方法的格式为float64
}

type rect struct {
	//struct类型是一种作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型rect，代表一个容器。这个容器拥有属性：也就是下面声明的width, height float64。
	//这样的类型称为struct。
	width, height float64 //将为 rect 和 circle 实现该接口。
}
type circle struct { //同上，创造一个包含 了 数据类型为float64的名为radius的函数 的 容器，这个容器也就是 一个struct数据类型的名为circle的函数。
	radius float64
}

func (r rect) area() float64 { //这里的(r rect)就是var r rect ，也就是声明r就是rect。
	//等效于func (rect).area() float64{}
	//这里就相当于调用了容器rect和接口area()
	//接口定义了对象的行为，接口指定类型应具有的方法，类型决定如何实现这些方法。(rect).area()就是方法
	//要在 Go 中实现一个接口，我们只需要实现接口中的所有方法。
	//rect是上面声明的容器。
	//名为geometry的interface（接口）中包含了area() float64
	//(rect).area()就是调用reach里的函数area() float64
	//这里我们为 rect 实现了 geometry 接口，因为实现了这个接口里的所有方法。
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	//如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。 这儿有一个通用的 measure 函数，我们可以通过它来使用所有的 geometry。
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r) //结构体类型 circle 和 rect 都实现了 geometry 接口， 所以我们可以将其实例作为 measure 的参数。
	measure(c)
}

```

**输出：**

```go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```



一个关于func的复习

普通函数声明（定义）



函数声明包括函数名、形式参数列表、返回值列表（可省略）以及函数体。

func 函数名(形式参数列表)(返回值列表){undefined
    函数体
}



形式参数列表描述了函数的参数名以及参数类型，这些参数作为局部变量，其值由参数调用者提供，返回值列表描述了函数返回值的变量名以及类型，如果函数返回一个无名变量或者没有返回值，返回值列表的括号是可以省略的。

如果一个函数声明不包括返回值列表，那么函数体执行完毕后，不会返回任何值