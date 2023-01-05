package main

import "fmt"

/*
接口是两件事：它是一组方法，但它也是一种类型。
*/
//设置一个interface，包括返回一个string类型的返回值的方法Speak()
type Animal interface {
	Speak() string
}

// 定义一个struct Dog，依附于struct Dog的方法Speak()可以实现interface Animal中的所有方法，所以struct Dog可以实现interface Animal
type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	CatMean string
}

func (c Cat) Speak() string {
	return "Meow!" + c.CatMean
}

// 相当于struct的方法给函数分组，用方法结合interface的输入值给函数分组
func UseInterface(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	animals := []Animal{Dog{}, Cat{}} //此处是interface{}作为一个 “空接口” 类型。[]interface{} 是一个slice（切片）可以存储任何类型的值
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
	// 实例化Cat struct
	catReal := Cat{
		CatMean: "catcatcat",
	}
	//因为cat struct实现了interface，所以输入值为interface的函数可以直接填写cat struct
	UseInterface(catReal)
}
