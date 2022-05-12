# go-16-055-Break 与 continue

# break

## 可使用`break`退出循环。

可以先写一个`for`结构的循环，再用`if`语句里面加上用于退出运行的`break`，就能达到具有条件判断功能，可自己停止的循环。

`switch`结构也能用`break`退出。

```go
package main

import (
	"fmt"
)

//可以在这里赋值，var i = 5

func main() {
	var i = 5 //也可以在这里赋值
	for {
		//如果在for语句的代码块内赋值，就会导致每次减一之后重新被赋值，无限死循环了。
		i = i - 1
		fmt.Println(i)
		if i < 0 { //每次迭代检查一遍条件，满足为止
			break
		}
	}
}

```

输出：

```
4
3
2
1
0
-1
```

## 嵌套的循环体，break 只会退出最内层的循环：

放在最内侧的循环里面，达到退出整个嵌套结构的目的。

一个例子：

```go
package main

func main() {
	for i:=0; i<3; i++ {
		for j:=0; j<10; j++ {
			if j>5 { //将此处改为 6 则 输出0123456  0123456  0123456  ，循环几次由外层决定，内层决定从零到几。
			    break   
			}
			print(j)
		}
		print("  ")
	}
}
```

输出：

	012345 012345 012345

把`if`后面的判断条件从`j`换成`i`的话，就是只影响循环几次。

```GO
package main

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 10; j++ {
			if i >= 3 {
				break
			}
			print(j)
		}
		print("  ")
	}
}
```

输出：

```
0123456789  0123456789  0123456789  
```

# continue

continue就像是特殊的goto语法，和`if`语句相结合之后，为他的发挥设定条件，若判断条件通过则跳过此次历遍。

```go
package main

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		print(i)
		print(" ")
	}
}
```

输出：

```
0 1 2 3 4 6 7 8 9
```

显然，5 被跳过了。

 

