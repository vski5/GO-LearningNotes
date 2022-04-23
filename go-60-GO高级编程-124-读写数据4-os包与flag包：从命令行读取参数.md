# 从命令行读取参数

相关阅读：[[go-57-GO高级编程-120-121-读写数据1-读取用户的输入]]
目前，Go 语言中使用最广泛功能最强大的命令行解析库是 cobra，但丰富的功能让 cobra 相比标准库的 flag 而言，变得异常复杂，为了减少使用的复杂度，cobra 甚至提供了代码生成的功能，可以自动生成命令行的骨架。然而，自动生成在节省了开发时间的同时，也让代码变得不够直观。
# os包
回忆一下如何给切片定义：`arr4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}	`
一个例子：
```go
// os_args.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")//从切片第二个元素开始加，因为第一个元素是程序本身的名字。
	}
	fmt.Println("Good Morning", who)
}

```

类型为string的切片变量`os.Args`，大概长这样：`[...]string{z,b,c}`

这个函数除了默认的`go run`以外还有特别的手法启动:
`go run main.go 参数1 参数2 参数3`

程序启动后`os.Args`会读取命令行输入的参数，这些命令行中的参数以空格分隔。放置在切片 `os.Args[]` 中。

`os.Args[0]` 放的是程序本身的名字，在本例中是 `os_args`。

# flag包
flag包有个功能，可以用来解析命令行选项。
但这个功能经常被用来替换基本常量。
可以在命令行给常量一些不一样的值。

前情提要：
1. flag包里的Flag结构体：
```go
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}
```

2. echo是一个计算机命令，

如果基于TCP协议，服务器就在TCP端口7检测有无消息，
如果使用UDP协议，基本过程和TCP一样，检测的端口也是7。 

echo是路由也是网络中最常用的数据包，可以通过发送echo包知道当前的连接节点有那些路径，并且通过往返时间能得出路径长度。

下面的程序 `echo.go` 模拟了 Unix 的 echo 功能：
```go
package main

import (
	"flag" // command line option parser
	"os"
)

//func flag.Bool(name string, value bool, usage string) *bool
//Bool定义了带有指定名称、默认值和使用字符串的Bool标志。
//返回值是存储标志值的bool变量的地址。
var NewLine = flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
//`flag.Bool()` 定义了一个默认值是 `false` 的 flag：当在命令行出现了第一个参数（这里是 "n"），flag 被设置成 `true`（NewLine 是 `*bool` 类型）。
const (
	Space   = " "
	Newline = "\n"
)

func main() {
	flag.PrintDefaults() //`flag.PrintDefaults()` 打印 flag 的使用帮助信息
	flag.Parse() //扫描参数列表（或者常量列表）并设置 flag
	var s string = ""
	for i := 0; i < flag.NArg(); i++ { //`flag.Narg()` 返回参数的数量。解析后 flag 或常量就可用了。 
		if i > 0 {
			s += " "
			//flag 被解引用到 `*NewLine`，所以当值是 `true` 时将添加一个 Newline（"\n"）。
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		/*
		 `flag.Arg(i)` 表示第i个参数。`Parse()` 之后 `flag.Arg(i)` 全部可用，
		 `flag.Arg(0)` 就是第一个真实的 flag，而不是像 `os.Args(0)` 放置程序的名字。
		*/
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

```

`flag.Parse()` 扫描参数列表（或者常量列表）并设置 flag, `flag.Arg(i)` 表示第i个参数。`Parse()` 之后 `flag.Arg(i)` 全部可用，`flag.Arg(0)` 就是第一个真实的 flag，而不是像 `os.Args(0)` 放置程序的名字。

`flag.Narg()` 返回参数的数量。解析后 flag 或常量就可用了。  
`flag.Bool()` 定义了一个默认值是 `false` 的 flag：当在命令行出现了第一个参数（这里是 "n"），flag 被设置成 `true`（NewLine 是 `*bool` 类型）。flag 被解引用到 `*NewLine`，所以当值是 `true` 时将添加一个 Newline（"\n"）。

`flag.PrintDefaults()` 打印 flag 的使用帮助信息，本例中打印的是：

```go
-n=false: print newline
```

`flag.VisitAll(fn func(*Flag))` 是另一个有用的功能：按照字典顺序遍历 flag，并且对每个标签调用 fn （参考 15.8 章的例子）

当在命令行（Windows）中执行：`echo.exe A B C`，将输出：`A B C`；执行 `echo.exe -n A B C`，将输出：

```
A
B
C
```

每个字符的输出都新起一行，每次都在输出的数据前面打印使用帮助信息：`-n=false: print newline`。

对于 `flag.Bool` 你可以设置布尔型 flag 来测试你的代码，例如定义一个 flag `processedFlag`:

```go
var processedFlag = flag.Bool("proc", false, "nothing processed yet")
```

在后面用如下代码来测试：

```go
if *processedFlag { // found flag -proc
	r = process()
}
```

要给 flag 定义其它类型，可以使用 `flag.Int()`，`flag.Float64()`，`flag.String()`。


# 阿里技术公众号的flag包中的 flag.Parse方法 解读


定义好命令行参数后，只需要调用 flag.Parse方法即可。
```go
package main  
  
import (  
   "flag"  
 "fmt")  
  
func main() {  
   var limit int  
 /*  
 func flag.IntVar(p *int, name string, value int, usage string) IntVar定义了一个具有指定名称、默认值和提示符的int标志。  
 参数p指向一个int变量，用于存储标志的值。  
 */ flag.IntVar(&limit, "limit", 9, "the max number of results")  
   flag.Parse() //这个函数会扫描命令行里的参数，赋值给上面声明的limit  
 fmt.Println("the limit is", limit)  
}
```

vscode执行就失败了，但是goland执行是ok的。

```go
// 执行结果
PS D:\golandproject\flagParse> go run flagParse.go
the limit is 10
//定义命令行参数之后
PS D:\golandproject\flagParse> go run flagParse.go -limit 100
the limit is 100

```

定要好命令行参数后，只需要调用 flag.Parse就可以实现参数的解析。

在定义命令行参数时，可以指定默认值以及对这个参数的使用说明。

# 阿里技术公众号的cobra 库 实现命令行读取：cobra可处理子命令
cobra可处理子命令。
子命令和跟命令遵循相同的定义模板，子命令还可以定义自己子命令

官方使用案例如下：
```go
package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var echoTimes int

	var cmdPrint = &cobra.Command{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long:  `print is for printing anything back to the screen.For many years people have printed back to the screen.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Print: " + strings.Join(args, " "))
		},
	}

	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long:  `echo is for echoing anything back.Echo works a lot like print, except it has a child command.`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	var cmdTimes = &cobra.Command{
		Use:   "times [string to echo]",
		Short: "Echo anything to the screen more times",
		Long:  `echo things multiple times back to the user by providing a count and a string.`,
		Args:  cobra.MinimumNArgs(1), //小于1一个值时返回错误
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}
	//cmdTimes是&cobra.Command（一个类型结构体的实际值）。
	//IntVarP定义了一个具有指定名称、默认值和提示符的int标志，但它接受一个可以在一个破折号(--)后使用的速记字母。
	//&echoTimes是用于存储标志的值。
	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	var rootCmd = &cobra.Command{Use: "app"} //声明rootCmd的结构体和&cobra.Command一样，但是换了use这部分。
	//AddCommand向父命令添加一个或多个命令。func (*cobra.Command).AddCommand(cmds ...*cobra.Command)
	rootCmd.AddCommand(cmdPrint, cmdEcho) //把子命令cmdPrint, cmdEcho加到父命令rootCmd
	cmdEcho.AddCommand(cmdTimes)         //把子命令cmdTimes加到父命令cmdEcho
	rootCmd.Execute()                    //Execute使用参数os.Args[1:]（默认情况下），并遍历命令树，为命令找到适当的匹配，然后是相应的标志。
}


```

有一个雷点：一定要带参数，还是符合规范的参数，不然就会报错跑不出来。

输出
>$ go run cobra.go echo times hello --times 3
Echo: hello
Echo: hello
Echo: hello

这里的使用方法还是很意思的。在上面的例子里分了3种用法，分别由三个结构体来定义。
>go run cobra.go echo times hello --times 3
>go run cobra.go print hello   
>go run cobra.go echo hello

其中`cmdTimes`是`cmdEcho`的子结构体，所以可以结合使用，`cmdEcho`和`cmdTimes`就是表明我要按这个顺序开始干活儿了。
但是没有给`cmdPrint`定义子结构体所以只能用他自己。

这个命令是在`Use: "print [string to print]",` 这一栏定义的吗？
测试一下：
先把原文件里的print改成text
>$ go run cobra.go text hello 
直接报错

是在cmdPrint里定义的吗？先把cmdPrint全部替换成cmdText
>$ go run cobra.go text hello 
>直接报错

把所有print换成text都是用go run cobra.go print hello  来操作。
我麻了。

把其他部分删光才知道，go run cobra.go print hello 里有用的是go run cobra.go，后面的都会被加上去，print可有可无。

恢复完全体的代码，又开始认 print 了。把 print换成text就要go run cobra.go text hello来操作。具体什么原因，我尚且在蒙古里。
go run cobra.go app --help
go run cobra.go Text hello
go run cobra.go text2 hello

## 我找到了1
 `**cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input") `
**在本地分配一个标志，该标志仅适用于该特定命令。**
**"times"是参数的关键词，
在`go run cobra.go echo times hello --times 3`**里面`--times 3`这个`times`就是适用于特定命令的标志。
VarP系列的函数都能定义。
此处是为cmdTimes定义句柄（标志）。

**但这个只能用在子命令这一行！！！**
**但这个只能用在子命令这一行！！！**
**但这个只能用在子命令这一行！！！**

加上--才是flag的标志。
>PS D:\golandproject\cobra> go run cobra.go echo times hello times 3   
Echo: hello times 3
PS D:\golandproject\cobra> go run cobra.go echo times hello --times 3 
Echo: hello
Echo: hello
Echo: hello

## 我找到了2
起作用的是cobra.Command这个结构体。

cobra.Command这个结构体里的run定义是起作用的，是真的会执行的函数。
 
cobra.Command这个结构体里的use定义 是作为 父命令的定义，
>Use:   "prin [string to print]",
>改成
>Use:   "secondtest [string to print]",

操作方式就变成了：go run cobra.go secondtest hello



# 起作用的是cobra.Command这个结构体。这只是一部分，还有百八行没写出来。
```go
type Command struct {  
// Use is the one-line usage message.  
// Recommended syntax is as follow:  
// [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.  
// ... indicates that you can specify multiple values for the previous argument.  
// | indicates mutually exclusive information. You can use the argument to the left of the separator or the  
// argument to the right of the separator. You cannot use both arguments in a single use of the command.  
// { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are  
// optional, they are enclosed in brackets ([ ]).
}

//一个例子：
 var cmdPrint = &cobra.Command{ //cobra包里的command类型的结构体的实际值
 Use: "print [string to print]",
 Short: "Print anything to the screen",
 Long: `print is for printing anything back to the screen.For many years people have printed back to the screen.`,
 Args: cobra.MinimumNArgs(1),
 Run: func(cmd *cobra.Command, args []string) {
 fmt.Println("Print: " + strings.Join(args, " "))

 },

```
我英语也不好，翻译一下结构大概就是：
```go
type Command struct {  
// Use: "一行利用命令行界面执行的程序显示的简短信息"  
// 推荐的语法如下:
// []标识一个可选参数，没有括在中括号中的参数是必需的。
// ... 表明 可以为前一个参数指定多个值。  
// | 标识互斥的信息. 可以使用分隔符左侧的参数，也可以使用分隔符右侧的参数。您不能在命令的一次使用中同时使用两个参数。 
// { } 当需要其中一个参数时，分隔一组互斥参数。如果是可选的，它们被括在大括号中 ([ ]).
}

```
usage message=利用命令行界面执行的程序显示的简短信息



但cobra有两个缺点：
1. 参数定义跟命令逻辑分离
2. 子命令与父命令的顺序定义不够灵活

