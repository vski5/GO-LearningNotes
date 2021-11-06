# go-7-047-strings 和 strconv 包



##  `strings` 包 

##  可用于对字符串的主要操作



### 前缀  和  后缀

1. `HasPrefix` 判断字符串 `s` 是否以 `prefix` 开头：  

prefix意思是 字首、词头

```go
strings.HasPrefix(s, prefix string) bool
```

用汉语来翻译就是

```
strings.HasPrefix(s, 你希望检测的开头有的东西而且她的格式是string) bool
```



2. HasSuffix` 判断字符串 `s` 是否以 `suffix` 结尾：

suffix意思是 后缀

```
trings.HasSuffix(s, suffix string) bool
```

用汉语来翻译就是

```
strings.HasSuffix(s, 你希望检测的结尾有的东西而且她的格式是string) bool
```

举个例子：

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \" %s \" have prefix %s? ", str, "Th")
	// 这里的\"%s\是占位符，%s输出字符串表示string格式，后面的str就填充进来了。是格式控制
	//一些字母前加"\"来表示常见的那些不能显示的ASCII字符，如\0,\t,\n等，就称为转义字符，就像是\n的意思是换行，再加一个\就再次转换意思。
	//在%s两边的\" 保证了 " 不是作为代码的一部分而是一个表意思的符号。"本意是代码的一部分，在前面加了一个\就把 " 的意思转换了。从功能意义到字面上的意义。
	fmt.Printf("%t\n",strings.HasPrefix(str, "Th"))
} //这里的%t\n，前半截%t是布尔占位符，后半截\n是换行的意思。
//这里的代码不能分行，要在一行里写完，不然就会报错
//strings.HasPrefix(str, "Th")就是在判断此处str所表示的string类型的函数开头是不是Th，也就是逗号后面的引号里的东西。

```

输出：

	T/F? Does the string "This is an example of a string" have prefix Th? true



### 字符串包含关系

`Contains` 判断字符串 `s` 是否包含 `substr`：

```go
strings.Contains(s, substr string) bool
```



### 判断子字符串或字符在父字符串中出现的位置（索引）

1. **正着查你想查的第一个在哪**

`Index` 返回字符串 `str` 在字符串 `s` 中的索引

`Index` 开头的 I 要大写

（`str` 的第一个字符的索引），

-1 表示字符串 `s` 不包含字符串 `str`：

```go
strings.Index(s, str string) int
```



2. **反着查你想查的第一个在哪**

`LastIndex` 返回字符串 `str` 在字符串 `s` 中最后出现位置的索引

（`str` 的第一个字符的索引），

-1 表示字符串 `s` 不包含字符串 `str`：

```go
strings.LastIndex(s, str string) int
```

**<u>注：上面两个string指的是str的数据类型要为string</u>**



3. 如果需要查询非 ASCII 编码的字符在父字符串中的位置，建议使用以下函数来对字符进行定位：

```go
strings.IndexRune(s string, r rune) int
```

```
注: 原文为 "If ch is a non-ASCII character（特征、特性） use strings.IndexRune(s string, ch int) int."

该方法在最新版本的 Go 中定义为 func IndexRune(s string, r rune) int

实际使用中的第二个参数 rune 可以是 rune 或 int, 例如 strings.IndexRune("chicken", 99) 或 strings.IndexRune("chicken", rune('k'))
```





- byte 等同于int8，常用来处理ascii字符
- rune 等同于int32,常用来处理unicode或utf-8字符

字符串的内容（纯字节）可以通过标准索引法来获取，在中括号 `[]` 内写入索引，索引从 0 开始计数：

- 字符串 str 的第 1 个字节：`str[0]`
- 第 i 个字节：`str[i - 1]`
- 最后 1 个字节：`str[len(str)-1]`

这种转换方案只对纯 ASCII 码的字符串有效。

### 字符串替换

`Replace` 用于将字符串 `str` 中的前 `n` 个字符串 `old` 替换为字符串 `new`，并返回一个新的字符串，

如果 `n = -1` 则替换所有字符串 `old` 为字符串 `new`：

```go
strings.Replace(str, old, new, n) string
```

### 统计字符串出现次数

`Count` 用于计算字符串 `str` 在字符串 `s` 中出现的非重叠次数：

```go
strings.Count(s, str string) int
```

示例

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
}
```

输出：

	Number of H's in Hello, how is it going, Hugo? is: 2
	Number of double g’s in gggggggggg is: 5





### 重复字符串

`Repeat` 用于重复 `count` 次字符串 `s` 并返回一个新的字符串：



```go
strings.Repeat(s, count int) string		

```

也就是

```
strings.Repeat(s, 重复数据格式为int的数字) string
```

示例

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var origS string = "Hi there! "
	var newS string  //先要声明newS的数据类型，下面用上了。

	newS = strings.Repeat(origS, 3) //重复origS所代表的东西3次
	fmt.Printf("The new repeated string is: %s\n", newS)
}
```

输出：

	The new repeated string is: Hi there! Hi there! Hi there!



### 修改字符串大小写

`ToLower` 将字符串中的 Unicode 字符全部转换为相应的小写字符：

```go
strings.ToLower(s) string
```

`ToUpper` 将字符串中的 Unicode 字符全部转换为相应的大写字符：

```go
strings.ToUpper(s) string
```

示例 4.17 [toupper_lower.go](examples/chapter_4/toupper_lower.go)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig) //正常输出
	lower = strings.ToLower(orig) //把这些全小写
	fmt.Printf("The lowercase string is: %s\n", lower) //把刚刚全小写的打印出来
	upper = strings.ToUpper(orig)  //把这些全大写
	fmt.Printf("The uppercase string is: %s\n", upper)  //把刚刚全大写的打印出来
}
```

输出：

	The original string is: Hey, how are you George?
	The lowercase string is: hey, how are you george?
	The uppercase string is: HEY, HOW ARE YOU GEORGE?





### 修剪字符串

使用 `strings.TrimSpace(s)` 来剔除字符串`s`开头和结尾的空白符号；

剔除指定字符，使用 `strings.Trim(s, "想剔除的字符")` 来将开头和结尾的 `像剔除的字符` 去除掉。

该函数的第二个参数可以包含任何字符

只除去开头的就用`strings.TrimLeft(s, "想剔除的字符")` 

只除去结尾的就用`strings.TrimRight(s, "想剔除的字符")` 

**剔除的是字符串`s`里的**



### 分割字符串

`strings.Fields(s)` 将会利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，

并返回一个 slice，如果字符串只包含空白符号，则返回一个长度为 0 的 slice。

`strings.Split(s, sep)` 用于自定义分割符号来对指定字符串进行分割，同样返回 slice。

`sep`就是用来切割的符号



### 拼接 slice 到字符串

**首先要知道如何理解** `for := range`

**`for 用来计数的索引 你需要遍历(就是每一个拆出来算)的字符串 := range`**

```go
package main

import (
	"fmt"
)

func main() {
	var str = "123456"
	for key, value := range str {
		//先把str的值赋给key和value，然后下面会直接用key和value来运算
		//for range 的组合，对字符串进行遍历，遍历(就是每一个拆出来算)时，key代表字符串的索引（base0）， value代表字符串中的每一个字符。
		//用下面两个可以用来格式化的占位符分别运算。
		//for ABC := range 就是依次代入运算A,B,C。一个个的算。但输出的key用来计数

		fmt.Printf("key:%d value:0x%x\n", key, value)
		//%d是给key留位置的占位符。表示要把被留位置的格式化成十进制，但输出的key用来计数，%d就是用十进制计数。
		//例如Printf("%d", 0x12)，输出得18，这里是把十八进制的0x12转换为了十进制的18.

		//%x是名为value的变量 的占位符。表示表示要把被留位置的格式化成十六进制，
		//例如Printf("%x", 13)，输出得d，这里是把十进制的13转换为了十六进制的d.

		//代码中的value变量，实际类型是rune，实际上就是 int32，以十六进制打印出来就是字符的编码。
	}
}

```



`Join` 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串：

```go
strings.Join(sl []string, sep string) string
```

示例

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog" //这句话是一个英语梗，刚好包括了26个字母
	sl := strings.Fields(str) //strings.Fields也就给后面的字母加了个[]
	fmt.Println(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val) 
		//此处的"%s - "里的%s就是挨个输出val里的东西，后面的 - 可以使输出结果每一个间隔一个 -
	}
	 //for 用来计数的索引 你需要遍历(就是每一个拆出来算)的字符串 := range
    //for := range 在此处下划线 让在 for := range 中 用来计数的第一项被舍弃，把val用花括号里面的计算方式，也就是函数经行运算，复习一下:= 是一种简明的赋值语句。
	fmt.Println()//用来打印上面的所有或者第一项？
	//上面的func main() 是声明main包，这里是指调用了包外的函数fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|") 
	//但这里的strings.Split也就给后面的字母加了个[]
	//strings.Split(s, sep)用于自定义分割符号来对指定字符串进行分割，同样返回 slice
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {      //用来计数的第一项被舍去
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)

}
```

输出：

	Splitted in slice: [The quick brown fox jumps over the lazy dog]
	The - quick - brown - fox - jumps - over - the - lazy - dog -
	Splitted in slice: [GO1 The ABC of Go 25]
	GO1 - The ABC of Go - 25 -
	sl2 joined by ;: GO1;The ABC of Go;25



## 4.7.11 从字符串中读取内容

函数 `strings.NewReader(str)` 用于生成一个 `Reader` 并读取字符串中的内容，

然后返回指向该 `Reader` 的指针，

说白了就是生成一个读取了`str`的值，如何指定一个东西指针定向到读取的位置。

从其它类型读取内容的函数还有：

- `Read()` 从 []byte 中读取内容。
- `ReadByte()`  从字符串中读取下一个 byte。
- 和 `ReadRune()` 从字符串中读取下一个rune。



## 4.7.12 字符串与其它类型的转换

`strconv`包实现字符串相关的类型转换。

类型T表示泛型，就是说所有类型、任何类型。

类型T都能转换为字符串。

`strconv`包里面有函数可以用来获取，你这个程序运行的操作系统下的int类型所占的位数，`strconv.IntSize`可以用来实现这个要求。



**Go提供了将数字类型转换为字符串的函数，如下所示：**

 1. `stronv.Itoa（i int） string`可以返回数字i所表示的类型的十进制的数，此处的(i int)指int类型的名为i的值

    例如`newS = strconv.Itoa(an)` //此处就是在算an所表示的类型的十进制的表示，并将其赋值到newS里。

 2.  `strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string` 

    单词Format是格式的意思。

     这个函数可以将64 位浮点型的数字转换为字符串，

    其中 `fmt` 表示格式（其值可以是 `'b'`、`'e'`、`'f'` 或 `'g'`），

    `prec` 表示精度，

    `bitSize` 则使用 32 表示 float32，用 64 表示 float64。

     `strconv.FormatFloat(64位的值, byte(字符)类型的格式, int格式表示的精度值, 此处写32或者64-分别表示float32和float64-要为int格式) string` 

**字符串有时候不能转化为其他的值**



**字符串类型转化为数字类型的函数：**

1. `strconv.Atoi(s string) (i int, err error)` 将字符串转换为 int 型。
2. `strconv.ParseFloat(s string, bitSize int) (f float64, err error)` 将字符串转换为 float64 型。

利用多返回值的特性，这些函数会返回 2 个值，

第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误，

因此，我们一般使用以下形式来进行从字符串到其它类型的转换：

	val, err = strconv.Atoi(s)

一个例子：

示例 4.19 [string_conversion.go](examples/chapter_4/string_conversion.go)

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string //上面三个先声明了会用到的值的类型

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize) //%d是十进制。strconv.IntSize用来获取：你这个程序运行的操作系统下的int类型所占的位数

	an, _ = strconv.Atoi(orig) //仅仅需要返回值的一部分的话，可以使用空白标识符 _ 。 _ 直接抛弃了可能出现的错误的返回值
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5             //新的an等于旧的an加5
	newS = strconv.Itoa(an) //此处就是在算an所表示的类型的十进制的表示，并将其赋值到newS里。
	fmt.Printf("The new string is: %s\n", newS)
}

```

输出：
	

	64 位系统：
	The size of ints is: 64
	32 位系统：
	The size of ints is: 32
	The integer is: 666
	The new string is: 671

