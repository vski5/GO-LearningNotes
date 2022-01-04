# go-15-054-for结构和goto语句

for就是一种重复执行某些语句的结构，在go里面只有for能做到这一点。

## 基本结构

	for 初始化语句; 条件语句; 修饰语句 {被重复执行的语句}

例如：

```go
for i := 0; -10< i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
```

输出：

```
This is the 0 iteration
This is the 1 iteration
This is the 2 iteration
This is the 3 iteration
This is the 4 iteration
```

这里的输出是从0开始的，因为设定的是`i := 0`;

如果设定的是`i := 3`，那就是从3开始执行，执行到满足 `i < 5`这个条件未知，也就是：

会输出:

```
This is the 3 iteration
This is the 4 iteration
```

条件语句 `i < 5;`，在每次循环开始前都会进行判断，一旦判断结果为 false，则退出循环体。

最后一部分为修饰语句 `i++`，一般用于增加或减少计数器，

`i--`就是反向数数，也就是反着数，例如

```go
func main() {
	for i := 3; 1 < i; i-- {
		fmt.Printf("This is the %d iteration\n", i)
	}
}
```

就是从5数到2，也就是输出：

```
This is the 5 iteration
This is the 4 iteration
This is the 3 iteration
This is the 2 iteration
```



这些组成部分，之间使用分号 `;` 相隔。

## 格式要求——

1. 左花括号 `{` 必须和 for 语句在同一行，
2. 计数器的生命周期在遇到右花括号 `}` 时便终止。
3. 一般习惯使用 i、j、z 或 ix 等较短的名称命名计数器。

## 可以在循环体中使用多个计数器,多个条件

**<u>两个分号`;`之间的多个条件用逗号`,`隔开</u>**

例如：

```
for i, j := 0, N; i < j; i, j = i+1, j-1 {}
```

## 可以将两个 for 循环嵌套起来

```go
for i:=0; i<5; i++ {
	for j:=0; j<10; j++ {
		println(j)
	}
}
```

## 一个花活儿：用 for结构 循环迭代一个Unicode编码的字符串

```go
package main

import "fmt"

func main() {
	str := "beautiful"
	fmt.Printf("str的长度是: %d\n", len(str))
	for ix :=0; ix < len(str); ix++ { //ix < len(str)意思是输出比str字节长度小的ix
		fmt.Printf("第 %d 个符号 的位置是: %c \n", ix, str[ix])
	}
	str2 := "中文"
	fmt.Printf("str2的长度是: %d\n", len(str2))
	for ix :=0; ix < len(str2); ix++ {
		fmt.Printf("第 %d 个符号 的位置是: %c \n", ix, str2[ix])
	}
}
```

会输出：

```
str的长度是: 9
第 0 个符号 的位置是: b 
第 1 个符号 的位置是: e 
第 2 个符号 的位置是: a 
第 3 个符号 的位置是: u 
第 4 个符号 的位置是: t 
第 5 个符号 的位置是: i 
第 6 个符号 的位置是: f 
第 7 个符号 的位置是: u 
第 8 个符号 的位置是: l 
str2的长度是: 6
第 0 个符号 的位置是: ä 
第 1 个符号 的位置是: ¸ 
第 2 个符号 的位置是: ­ 
第 3 个符号 的位置是: æ 
第 4 个符号 的位置是:  
第 5 个符号 的位置是:  

```

**第二部分，中文用的是Unicode编码，GO把“中文”拆成单个字节就会出问题，<u>所以不能用索引判断是否为相同的字符</u>。**



# goto语句

## goto语句的基本结构

可以网上找一下逻辑图。



一个例子

```
package main

import "fmt"

func main() {
   var a int = 10 //定义局部变量
//下面开始写循环内容
   LOOP: for a < 20 {
      if a == 15 {
         a = a + 1 //a+1加到a=15的时候，就goto LOOP，也就是跳到下一块，也就是舍去15。
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++     //舍去a++就会出现循环输出a的值为 : 10，然后就死机了，改成a--就会反向输出，一直到负的很大的数，然后就死机了。
   }  
}
```

**<u>//舍去a++就会出现循环输出a的值为 : 10，然后就死机了，改成a--就会反向输出，一直到负的很大的数，然后就死机了。</u>**

# 关于本章的题目

## 一个嵌套结构的实现

创建一个程序，要求能够打印类似下面的结果（尾行达 25 个字符为止）：

	G
	GG
	GGG
	GGGG
	GGGGG
	GGGGGG

1. 使用 2 层嵌套 for 循环。
2. 仅用 1 层 for 循环以及字符串连接。



两种实现方法，第一种是嵌套

```go
package main

func main() {
	// 1 - use 2 nested for loops
    //在这个历遍中，先判断第一层壳子符不符合要求，再判断第二层内核符不符合要求，让我觉得很烦的内核判断条件有外壳里出现的参数，虽然说到底是二元一次不等式方程组，我依旧觉得很烦。
    //例如，在第二个循环里面，会运行两遍里面的结构
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")   //这里就没有输出j，而是将j简单作为一个记数工具。
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
```

## 题目：使用按位补码从 0 到 10，使用位表达式 `%b` 来格式化输出。

解答：

```go
package main

import (
	"fmt"
)

func main() {

	for i := 0; i <= 10; i++ {
		fmt.Printf("%b , %b \n ", i, ^i)
	}
}
```

#### 知识点回顾：（可见go-5-045）

**位异**  （^）参加运算的两个对象，如果两个相应位为“异”（值不同），则该位结果为1，否则为0。判断是否相同的符号。不同则答 是 ，相同则答 否

```
1 ^ 1 -> 0   // 位异 参加运算的两个对象，如果两个相应位为“异”（值不同），则该位结果为1，否则为0。
//判断是否相同的符号。不同则答 是 ，相同则答 否
1 ^ 0 -> 1
0 ^ 1 -> 1
0 ^ 0 -> 0
```

- 按位补足  `^`：

  与 异（用来判断是否不同的运算符） 或 运算符一同使用，

  即格式为 `m^x`时

  `^10` 里的^相当于一个占位符号，因为位运算只能算相同位数的，用^使数值位数变得一样。
  在默认情况下（无符号 x 使用，“全部位设置为 1”），相当于使后面的数值乘以一个 1 ，也就是`^10` 里的`^`前面多了个1，想到于上面提到的格式 `m^x` 之中的m=1。

  当x为负数时， `m^x` 之中的m=-1

  主要看这个例子。

  每一位都在每一位上加，就像是列了个加法竖式，十位和十位加，个位与各位加。

  但是是在二进制位上操作运算。

  	^10 = -01 ^ 10 = -11


##  Fizz-Buzz 问题

写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 `Buzz` 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 `FizzBuzz`（提示：使用 switch 语句）。

解答：

这里要用到`%`，表示整除后的余数

```go

package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {
	for i := 0; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
```

做这道题最大的问题就是没记清楚switch的结构，还有就是 for 后面要用`:=`带有冒号的赋值



###  使用 `*` 符号打印宽为 20，高为 10 的矩形。

```GO
package main

import "fmt"

func main() {
	w, h := 20, 10
	for y := 0; y < h; y++ { //打印9行
		for x := 0; x < w; x++ { //每行打印19个
			fmt.Print("*")
		}
		fmt.Println() //用来换行的
	}
}

```

# 基于条件判断的迭代

基本形式为：`for 条件语句 {}`。

```
package main

import "fmt"

func main() {
	var i int = 5

	for i >= 0 {
		i = i - 1
		fmt.Printf("The variable i is now: %d\n", i)
	}
}
```

### 无限循环

如果 for 循环的头部没有条件语句，那么就会认为条件永远为 true，因此循环体内必须有相关的条件判断以确保会在某个时刻退出循环。

退出循环的方法有二 

1. break语句，只退出当前的循环体
2. return语句，不执行语句后续代码，提前对函数返回

无限循环的应用：

服务器不断等待接受新的请求。

```go
for t, err = p.Token(); err == nil; t, err = p.Token() {
	...
}
```

# for-range 结构

语法上很类似其它语言中 foreach 语句

可以获得每次迭代所对应的索引。

一般格式为：

```go
for ix, val := range coll { }
ix 是数组或者切片的索引,
val是在该索引位置的值
大概会输出 
1 one
2 two 
.......
类似这样的格式的值
```

```go
package main

import "fmt"

func main() {
	str := "beautiful"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str { //这里是历遍打印出来返回值，str这个字符串里的东西就被挨个拆分，拆成一个个单独的字母。
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
    
	fmt.Println()
    
	str2 := "汉语"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 { 
    	fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
    
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
    
	for index, rune := range str2 {
    	fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
```

输出：

```
The length of str is: 9
Character on position 0 is: b 
Character on position 1 is: e 
Character on position 2 is: a 
Character on position 3 is: u 
Character on position 4 is: t 
Character on position 5 is: i 
Character on position 6 is: f 
Character on position 7 is: u 
Character on position 8 is: l 

The length of str2 is: 6
character 汉 starts at byte position 0
character 语 starts at byte position 3

index int(rune) rune    char bytes
0       27721      U+6C49 '汉' E6 B1 89
3       35821      U+8BED '语' E8 AF AD

```







# 关于println()和Printf()的观察

前者就像是python里的`print("")`，是在直接输出字符。

后者就需要加上`%d`这种格式化的语句，专业说法叫占位符，再一一对应才行，举个例子

```
var AAA,CCC str := go,python
fmt.Printf("learn what? , %c , %c ",AAA,CCC)
```

输出的就是

```
learn what? , go , python
```

