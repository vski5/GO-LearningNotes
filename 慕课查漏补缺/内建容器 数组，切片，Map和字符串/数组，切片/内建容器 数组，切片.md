# 内建容器 数组，切片，Map和字符串



## 数组

类似 ：

有三种表达方式

```go
package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int         //此处中括号为有5个int
	arr2 := [3]int{1, 3, 5} //此处中括号指有3个int，分别为1，3，5
	arr3 := [...]int{2, 4, 6, 8, 10}//[...]就是让编译器帮忙数有几个int
	var grid [2][3]int//2个长度为3的int数组
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
}
```

输出结果为：

```
[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
[[0 0 0] [0 0 0]]
```

可见未写数组内容的直接打印出 [ ] 内的数字的数量的0

`var grid [2][3]int`这种类型的就是两行三列。

- 数量写在类型前面

还是以上面的代码为例：

3个历遍arr3的方法

```go
package main

import (
	"fmt"
)

func main() {
	arr3 := [...]int{2, 4, 6, 8, 10}
	for k := 0; k < len(arr3); k++ { //这里是分号；隔开
		fmt.Println(arr3[k]) //这里的arr3[k]就是依次找出来arr3这个数组里面的第k个数字。
	}
	for i := range arr3 {
		fmt.Println(arr3[i]) //for i := range 经典历遍结构，arr3[i]就是调用数组第i个。
	}
	for f, k := range arr3 {
		fmt.Println(f, k) //历遍结构，此处f为序数，k为第f个元素对应的值。
	}

}

```

输出：

```
2
4
6
8
10
2
4
6
8
10
0 2
1 4
2 6
3 8
4 10
```

数组是<u>值类型</u>

值类型就是拷贝，必须规定长度，[ ]内不写就是切片而不是数组。

**[ ]内的数字不同，也就是int长度不同，就会被认为不是同一类型。**

```go
package main

import (
	"fmt"
)

func PrintArry(arr [5]int) { //arr为命名输入值，此处的[5]int是一种数值类型
	arr[0] = 99 //此处就是规定使用这个生造出来的函数的第0个（也就是第1个）都会变成99，只用了一个等于号
	for f, k := range arr {
		fmt.Println(f, k) //历遍结构，此处f为序数，k为依次的历遍的按顺序数出来的数字。
	}
}
func main() {

	var arr1 [5]int //此处中括号为有5个int

	arr2 := [3]int{1, 3, 5} //此处中括号指有3个int，分别为1，3，5

	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr2)

	PrintArry(arr1) //这里自己生造出来的函数PrintArry()，输入值要满足[5]int这种数值类型，arr2只有3个int显然不满足
	PrintArry(arr3)
	fmt.Println(arr1, arr3)

}

```

 输出：

```
[1 3 5]
0 99
1 0
2 0
3 0
4 0
0 99
1 4
2 6
3 8
4 10
[0 0 0 0 0] [2 4 6 8 10]
```

调用`func f(arr [10]int)`会 <u>拷贝</u> 数组，因为数组是值传递，只要是值传递就会走拷贝。

想改原值就用指针：

可以在**指针类型**前面加上 * 号（前缀）来**获取指针所指向的内容**，就像是此处的 *int 来获取指向int的内容。

变量前加一个&就能知道这个变量在内存的哪



这是GOlang的特色。

一般不直接用数组，用切片。

## 切片(Slice)

数组要明确有多少个，切片不用。

数组：

​	`arr := [...]int{1,2,3,4,5}` 

​	或者 `arr := [5]int{1,2,3,4,5}` 

切片Slice：

​	首先定义一个数组`arr := [...]int{1,2,3,4,5,6,7}` ，再写`s := arr[2:6]`  (读作2到6)，这个s就是切片，值就是[2 3 4 5]，因为半开半闭区间，左边的2包含进去，右边的6就不包括。

​	

```go
package main

import "fmt"

func UpdateSlice(s []int) { //此处方括号内没写长度，s就是一个int的slice
	s[0] = 100
}
func main() {
	arr4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr4[2:6]
	fmt.Println(s)
	fmt.Println(arr4[2:6]) //这四个arr4定义的方括号下标区间都是一种slice
	fmt.Println(arr4[:6])  //slice不是值类型，是一种视图view，此处是看最开头到6这个数。
	fmt.Println(arr4[2:])
	fmt.Println(arr4[:])
	s1 := arr4[2:]
	fmt.Println("s1=", s1)
	s2 := arr4[:]
	fmt.Println("s2=", s2)

	UpdateSlice(s1)
	//s1 := arr4[2:]，从2这个数开始，因为2是s1这个切片的第一个数，UpdateSlice(s1) 就是改第一个数为100，所以2就被改成100了
	fmt.Println("改过的arr4=", arr4)            // [0 1 100 3 4 5 6 7]
	fmt.Println("被UpdateSlice(s1)后的s1=", s1) // [100 3 4 5 6 7]
	//此处s1是对上面的view,所以是从上面第三个开始。
	UpdateSlice(s2) //s2 := arr4[:]， 完全的arr4
	//这里的UpdateSlice(s2) 就是改第一个数为100，所以0就被改成100了
	fmt.Println("s2=", s2)
	fmt.Println("因为这里的s2是上面改过的arr4的view，所以2变成了100。view读的是上面一个arr4，不是第一个")
	fmt.Println("arr4=", arr4)
}

```

输出：

```
[2 3 4 5]
[2 3 4 5]
[0 1 2 3 4 5]
[2 3 4 5 6 7]
[0 1 2 3 4 5 6 7]
改过的arr4= [0 1 100 3 4 5 6 7]
被UpdateSlice(s1)后的s1= [100 3 4 5 6 7]
s2= [100 1 100 3 4 5 6 7]
因为这里的s2是上面改过的arr4的view，所以2变成了100。view读的是上面一个arr4，不是第一个
arr4= [100 1 100 3 4 5 6 7]
```

slice本身没有数据，是对底层array的一个view（视图），能改变底层的**<u>array</u>**

类似  `s2 := arr4[:]`这种，就是数组取得slice的方法，这样就不用指针。

slice是从0开始数，第零个是多少，然后第一个是多少,

比如说arr5 := [...]int{9, 1, 2, 3, 4, 5, 6, 7}，这里的第零个就是9。

例如：

```go
	arr5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	x := arr5[2:6]
	fmt.Println(x)
	//还可以对slice取slice，这被称作re slice
	x = x[:5] //从第一个切片开始向右数5个
	fmt.Println(x)
	x = x[:6] //从第一个切片开始向右数六个
	fmt.Println(x)
	x = x[2:]
	fmt.Println(x) //但是这里就变成了切片的切片，被进一步裁剪。从零开始数，0.1.2然后就数到了4这个数。
	x = x[2:]
	fmt.Println(x) //但是这里就变成了切片的切片，被进一步裁剪。从零开始数，0.1.2然后就数到了6这个数。这个问题会在下面的“课后问题：slice的扩展了解”中讲到
```

输出

```
[2 3 4 5]
[2 3 4 5 6]
[2 3 4 5 6 7]
[4 5 6 7]
[6 7]

```



对数组其中某个值的赋值直接用=号

例如：`arr[0],arr[2] = 0,2`,这里就是对第0个数和第2个数赋值。



### 课后问题：slice的扩展了解

**记住：都是从0开始数的数**，**左开右闭区间**

```go
arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6]
s2 := s1[3:5]

```

输出的结果：

```
s1 = {2, 3, 4, 5}
s2 = {5,6} 
```

直接看上去s1根本就没有5和6，下面开始讲解为什么是这样。

本质上是对底层的arr的view而不是对s1的view，slice是可以扩展的。

![slice的view的结构](https://github.com/vski5/GO-LearningNotes/blob/2021-11-6/%E6%85%95%E8%AF%BE%E6%9F%A5%E6%BC%8F%E8%A1%A5%E7%BC%BA/%E5%86%85%E5%BB%BA%E5%AE%B9%E5%99%A8%20%E6%95%B0%E7%BB%84%EF%BC%8C%E5%88%87%E7%89%87%EF%BC%8CMap%E5%92%8C%E5%AD%97%E7%AC%A6%E4%B8%B2/slice%E7%9A%84view%E7%9A%84%E7%BB%93%E6%9E%84.jpg)



slice的实现：ptr指向开头的元素，len说明slice的长度，cap决定下面整个arr的长度就是最长能取的长度。

![slice的结构与实现](https://github.com/vski5/GO-LearningNotes/blob/2021-11-6/%E6%85%95%E8%AF%BE%E6%9F%A5%E6%BC%8F%E8%A1%A5%E7%BC%BA/%E5%86%85%E5%BB%BA%E5%AE%B9%E5%99%A8%20%E6%95%B0%E7%BB%84%EF%BC%8C%E5%88%87%E7%89%87%EF%BC%8CMap%E5%92%8C%E5%AD%97%E7%AC%A6%E4%B8%B2/slice%E7%9A%84%E7%BB%93%E6%9E%84%E4%B8%8E%E5%AE%9E%E7%8E%B0.jpg)

### lens和cap是可以取得的

```go
arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6]
s2 := s1[3:5]

fmt.Printf(s1,len(s1),cap(s1)) //这一步就是求s1和len和cap

```

