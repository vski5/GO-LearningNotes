# Map



map类型可以写为map[K]V，其中K和V分别对应key和value。
map中所有的key都有相同的类型，所有的value也有着相同的类型，
但是key和value之间可以是不同的数据类型。

**<u>string类型的key或者value要被双引号给括起来</u>**

## map有3种表示方式
1. 第一种是`n := map[key的类型]value的类型{"key1":value1,"key2":value2,"keyN":valueN,}`，

   第一种声明方法{}内的可以换行，具体如下。

2. 第二种是`m := make(map[key的类型]value的类型）`，

   然后在同一个代码块的这一行下面表明`m["key1"] = value1`，

   然后按回车另起一行写`m["key2"] = value2`直到第N个。

   不赋值就是一个空map

3. 第三种是`var m3 map[key]value`

   不赋值就是nil

```go
package main

import (
	"fmt"
)

func main() {
	//先写两种表示方式

	//第一种声明方式
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map-n:", n)
	//第一种声明方法{}内的可以换行，具体如下。
	Y := map[string]int{
		"alice":   31,
		"charlie": 34,
		"kitty":   39,
	}
	fmt.Println("map-Y:", Y)

	//第二种声明方式
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map-m:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m)) //记数，数m中有几个key/value对。

	delete(m, "k2") // 抛弃名为m的map中的key"K2"
	fmt.Println("map:", m)
	//当从一个 map 中取值时，还有可以选择是否接收的第二个返回值，该值表明了 map 中是否存在这个键。
	//这可以用来消除 键不存在 和 键的值为零值 产生的歧义， 例如 0 和 ""。
	//这里我们不需要值，所以用 空白标识符(blank identifier) _ 将其忽略。
	_, prsK1 := m["k1"] //ture 此处代码意思是用_舍去K1，K1健在，所以返回正确
	_, prsK2 := m["k2"] //false 此处代码意思是用_舍去K2，但K2被上面的delete（m,"K2")舍去，所以返回false
	//这个一般用于判断你写的KEY是否存在(也就是找m["K"]里的"K"是否存在)，如果不存在就会打印出来false
	prsK3, _ := m["k1"] //7 此处会报错：对空白标识符不必要的赋值,因为第二个值（K2）已经被舍去了。
	prsK4, _ := m["k2"] //0 同上。
	prsK5 := m["k1"]    //7
	prsK6 := m["k2"]    //0 key不存在则获取value类型的初始值
	fmt.Println("prs:", prsK1, prsK2, prsK3, prsK4, prsK5, prsK6)

	//历遍map的方法
	for kitty, alice := range Y { //此处for后面的值要有两个，且和下面括号内的要一致，同时要是range后面的map里的值，不然就会报错，顺序一样的话就会打出来K:Y,不一致就打出来Y:K
		fmt.Println(kitty, alice) //K/V的上下顺序是随机的，这是一个hash map。
	}
	//判断value是否存在的方法

	text1 := make(map[string]string)
	text1["apple"] = "banana"
	text1["blue"] = "yellow"
	fmt.Println(text1)

	if falseyellow, ok := text1["falseblue"]; ok {
		fmt.Println(falseyellow) //这里打印的结果是text1["blue"]所对应的value，不能检查value是否正确，只能检查key是否正确

	} else {
		fmt.Println("hello?")
	}
	//格式为 if value1 , ok := m["key"]; ok {
	//	fmt.Println(value1)
	//}else{
	//	fmt.Println("你想说的话")
	//}
}

```

输出：

```
map-n: map[bar:2 foo:1]
map-Y: map[alice:31 charlie:34 kitty:39]
map-m: map[k1:7 k2:13]
v1:  7
len: 2
map: map[k1:7]
prs: true false 7 0 7 0
alice 31
charlie 34
kitty 39
map[apple:banana blue:yellow]
hello?

```

默写练习：

```go
func main() { 
    m := map[key](value){
    "key1":value1
}
    n := make([key]value)
    n["key2"] = value2       
}

```

