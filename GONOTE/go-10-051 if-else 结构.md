# go-10-051 if-else 结构

if 用于条件判断，判断是否符合if后面跟的条件，

若成立则继续执行if后面的大括号内的代码块，

若不符合就跳过此处的代码。





```go
if condition {
	// do something	，condition是条件的意思
}
```

如果存在第二个分支，则可以在上面代码的基础上添加 else 关键字以及另一代码块，

这个代码块中的代码只有在条件不满足时才会执行。

if 和 else 后的两个代码块是相互独立的分支，只可能执行其中一个。

```go
if condition {
	// do something	
} else {
	// do something	
}
```



**如果存在第三个分支，则可以使用下面这种三个独立分支的形式：**

```go
if condition{
    
}else if condition2{
    
}else {
   
}
```

直接在后面叠加

```
else if condition{

}
```

就完事儿了，

如果是最后一个，那就是

```
else {
	// catch-all or default
}
```

else-if 分支的数量是没有限制的，但是为了代码的可读性，还是不要在 if 后面加入太多的 else-if 结构。

**代码敲出来的格式规范的问题**

1. 代码块外的{}不可省略，即使只有一行代码。

2. 关键字 if 和 else 之后的左大括号 `{` 必须和关键字在同一行

   ```go
   if condition1 {
   	// do something	
   } else if condition2 { //关键字 if 和 else 之后的左大括号 `{` 必须和关键字在同一行,上面的只有一个if也是和{在一行的，下面的else同理。
   	// do something else	
   } else {
   	// catch-all or default
   }
   ```

3. 如果你使用了 else-if 结构，则<u>**前段**代码块的右大括号 `}` 必须和 else-if 关键字在同一行</u>。

   ```go
   //一个例子
   if condition1 {
   	// do something	
   } else if condition2 { //前段代码块的右大括号 `}` 和 else-if 关键字在同一行
   	// do something else	
   } else {
   	// catch-all or default
   }
   
   ```

4. ​           2和3这两条规则都是被编译器强制规定的。

5. 在你使用 `gofmt` 格式化代码之后，每个分支内的代码都会缩进 4 个或 8 个空格，或者是 1 个 tab，并且右大括号与对应的 if 关键字垂直对齐。

   ```go
   package main
   import "fmt"
   func main() {
   	bool1 := true
   	if bool1 { //此处的if和下面的}对齐了
   		fmt.Printf("The value is true\n")
   	} else {   //此处的}和上面的if对齐了
   		fmt.Printf("The value is false\n")
   	}
   }
   
   
   ```

   ​	会输出

   ```
   The value is true
   ```

   **注意事项** 这里不需要使用 `if bool1 == true` 来判断，因为 `bool1` 本身已经是一个布尔类型的值。

这种做法一般都用在测试 `true` 或者有利条件时，

但你也可以使用取反 `!` 来判断值的相反结果，

如：`if !bool1` 或者 `if !(condition)`。

后者的括号大多数情况下是必须的，如这种情况：`if !(var1 == var2)`。

**一个BUG**

当 if 结构内有 `break`、`continue`、`goto` 或者 `return` 语句时，Go 代码的常见写法是省略 `else` 部分（另见第 5.2 节）。无论满足哪个条件都会返回 x 或者 y 时，一般使用以下写法：

```go
if condition {
	return x
}  //这个地方本来应该有一个else，现在省略了，下面再加上这个语句
return y 
```

**<u>这里举一些有用的例子</u>**：

1. 判断一个字符串是否为空：

   - `if str == "" { ... }`

   - ```go
     if str == ""{
     
     }else{
     
     }
     ```

     

   - `if len(str) == 0 {...}`	

   - ```go
     if len(str)==0{ //可以通过函数len()来获取字符串所占的字节长度，例如  len(str)
     
     }else{
     
     }
     ```

     

2. 判断运行 Go 程序的操作系统类型，这可以通过常量 `runtime.GOOS` 来判断(第 2.2 节)。

   	if runtime.GOOS == "windows"	 {
   		.	..
   	} else { // Unix-like
   		.	..
   	}

   这段代码一般被放在 init() 函数中执行。

   这儿还有一段示例来演示如何根据操作系统来决定输入结束的提示：

   ```go
   var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
   
   func init() {
   	if runtime.GOOS == "windows" {
   		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")		
   	} else { //Unix-like
   		prompt = fmt.Sprintf(prompt, "Ctrl+D")
   	}
   }
   ```

3. 函数 `Abs()` 用于返回一个整型数字的绝对值

```go
func Abs(x int) int {
	if x < 0 {
		return -x  //这里的x指的是上面那个x int 也就是int格式的数字x
	}
	return x	
}
```

4. `isGreater` 用于比较两个整型数字的大小:

```go
func isGreater(x, y int) bool { //x和y是两个int格式的值。后面的bool是声明isGreater(x, y int)的结果位bool型，要么对要么错
	if x > y {
		return true	
	}
	return false
}
```

在第四种情况中，if 可以包含一个 *初始化* 语句（如：给一个变量赋值）。

<u>**尽量不要这么做！！！**</u>

这种写法具有固定的格式——>在 *初始化* 语句后方必须加上 分号 ;

*Initialization*指 *初始化*，*初始化* 就是把变量赋为默认值

```go
if initialization; condition {
	// do something
}
```

例如:

```go
val := 10
if val > max {
	// do something
}
```

可以这样写:

```go
if val := 10; val > max {
	// do something
}
```

但使用简短方式 `:=` 声明的变量的作用域只存在于 if 结构中，也就是在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在。

再此种声明下，只有这个 if-else 结构内的代码块是声明的变量的作用域。

如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏。

最简单的解决方案就是不要在初始化语句中声明变量



示例 

```go
package main

import "fmt"

func main() {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Printf("first is less than or equal to 0\n")
	} else if first > 0 && first < 5 { //“&&”表示 与,意为同时都要满足。 “||”表示 或,意为二者或多着只要满足其中一个
		fmt.Printf("first is between 0 and 5\n")
	} else {
		fmt.Printf("first is 5 or greater\n")
	}
	if cond = 5; cond > 10 { 
		fmt.Printf("cond is greater than 10\n")
	} else {
		fmt.Printf("cond is not greater than 10\n")
	}
}
```

输出：

	first is 5 or greater
	cond is not greater than 10

下面的代码片段展示了如何通过在初始化语句中获取函数 `process()` 的返回值，并在条件语句中作为判定条件来决定是否执行 if 结构中的代码：

```go
if value := process(data); value > max {
	...
}
```
