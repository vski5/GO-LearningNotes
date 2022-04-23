上面学了 缓冲读取（bufio包）和 命令行输入参数（os.Args或者flag.Parse方法）。

以上两种都是不加参数的话，你输入什么屏幕就打印什么。

参数被认为是文件名，如果文件存在的话就打印文件内容到屏幕。

这种直接命令行加参数的手段就是CLI

命令行执行 `cat test` 测试输出。

# 用 buffer 读取文件，然后命令行执行 `cat test` 测试输出。
cat就Linux里面看日志的命令。

示例 12.11 cat.go：

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')     //读文件r每遇到源文件里的换行就输出，然后换行，直到EOF。
		fmt.Fprintf(os.Stdout, "%s", buf) //输出标准输出，后面的buf是会被标准输出的内容。
		if err == io.EOF {                //检测到EOF的时候就退出循环
			break
		}
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin)) //自造的cat函数输入值是为标准输入创造的缓冲空间。同时只有这样才会变成符合条件的类型。
	}
	//i会小于命令行输入的参数个数。
	for i := 0; i < flag.NArg(); i++ { //NArg是标记处理后剩余的参数数量。此处所指标记处理为flag.Parse()，获取命令行输入的参数。
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

```

在 12.6 章节，我们将看到如何使用缓冲写入。

**练习 12.6**：[cat_numbered.go](app://obsidian.md/exercises/chapter_12/cat_numbered.go)

扩展 cat.go 例子，使用 flag 添加一个选项，目的是为每一行头部加入一个行号。

使用 `cat -n test` 测试输出。->标志符号为n
循环需要判断所以用`var numberFlag = flag.Bool("n", false, "number each line") `


```go
package main  
  
import (  
   "bufio"  
 "flag" "fmt" "io" "os")  
//改动在此。制造一个默认值为false的，标志符号为n的flag.Bool  
var numberFlag = flag.Bool("n", false, "number each line")  
  
func cat(r *bufio.Reader) {  
   i := 1  
 for {  
      buf, err := r.ReadBytes('\n')     //读文件r每遇到源文件里的换行就输出，然后换行，直到EOF。  
 fmt.Fprintf(os.Stdout, "%s", buf) //输出标准输出，后面的buf是会被标准输出的内容。  
 if err == io.EOF {                //检测到EOF的时候就退出循环  
 break  
 }  
      //改动在此  
 if *numberFlag { //判断是否还有参数。  
         fmt.Fprintf(os.Stdout, "%5d %s", i, buf)  
         i++  
      } else {  
         fmt.Fprintf(os.Stdout, "%s", buf)  
      }  
   }  
   return  
}  
  
func main() {  
   flag.Parse()  
   if flag.NArg() == 0 {  
      cat(bufio.NewReader(os.Stdin)) //自造的cat函数输入值是为标准输入创造的缓冲空间。同时只有这样才会变成符合条件的类型。  
 }  
   //i会小于命令行输入的参数个数。  
 for i := 0; i < flag.NArg(); i++ { //NArg是标记处理后剩余的参数数量。此处所指标记处理为flag.Parse()，获取命令行输入的参数。  
 f, err := os.Open(flag.Arg(i))  
      if err != nil {  
         fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())  
         continue  
 }  
      cat(bufio.NewReader(f))  
      f.Close()  
   }  
}
```
