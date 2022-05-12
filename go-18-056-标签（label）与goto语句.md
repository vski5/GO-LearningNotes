# go-18-056-标签（label）与goto语句

for，switch，select三种结构都可以配合标签（label）形式的标识符使用，即某一行第一个冒号（`:`）结尾的单词。

gofmt 会将后续代码自动移至下一行

# label 标签

标签（label）的名称大小写敏感，为了可读性，一般用全大写，比如LABEL。

一个例子，

```go
package main

import "fmt"

func main() {

LABEL1: //此处相当于为LABEL1赋值
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL1 //符合if后的条件时，跳到上面的LABEL1:，继续开始i的递加，相当于跳过并列的的大箱子里面小箱子里的i==2以及后面的可能性。
                //（箱子理论可以见go-15-054里面的大标题for语句的嵌套执行顺序）
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

}
```

输出：

```
i is: 0, and j is: 0
i is: 0, and j is: 1
i is: 1, and j is: 0
i is: 1, and j is: 1
i is: 2, and j is: 0
i is: 2, and j is: 1
i is: 3, and j is: 0
i is: 3, and j is: 1

```



# goto语句的作用，goto语句和label的配合

一个例子：

```GO
package main

func main() {
	i := 0
HERE:
	print(i)
	i++ //指i在递加
	if i >= 6 {
		return //不执行语句后续代码，提前对函数返回
	}
	goto HERE //在i不大于或等于6的时候，返回标签HERE继续输出。
}

```

输出：

```
012345
```



**一个格式：**

在`if`语句后面的条件，等于号要写两个，和*大于或等于*是一个长度，就像是`==`,`<=`,`>=`。

**一个复习：**

退出循环的方法有二 

1. break语句，只退出当前的循环体
2. return语句，不执行语句后续代码，提前对函数返回

**一个格式要求：**

标签要在goto后面，不如说上面的`goto HERE`。

标签与格式之间不得有新的，对变量的定义，也就是不能有`var`，`:=`这类变量赋值语句。

**写在最后：一般不用标签和goto语句，因为会导致糟糕的可读性**
