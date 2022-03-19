# go-22-064-defer和追踪和Go的异常处理 defer, panic, recover

与`goto`的跳过满足要求的函数不同，

1. `defer`可以让我们推迟到函数返回之前一刻才执行某个语句或函数，
2. 或者是任意执行`return`的语句之后一刻才执行某个语句或函数

**<u>目的：</u>**`defer`用于确保程序在执行完成后，会调用某个函数，一般是执行清理工作。 Defer 的用途跟其他语言的 `ensure` 或 `finally` 类似。

## Go by Example 中文版中的Defer例子

这个例子的目的是演示一个 会去写入数据并在程序结束时自动关闭文件的代码，用`defer`做到。

```go
package main

import (
    "fmt"
    "os"
)

func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
```



### 多个defer

一个代码块中有多个defer文件时，会逆序执行，也就是类似栈的后进先出：

```go
func main(){
    for i:= 0 ;i<=5 ;i++{
        defer fmt.Println(i)
    }
}
```

会输出：

```
543210
```

### 一些关于defer的收尾技巧。

1. 关闭文件流

   ```go
   defer file.Close()
   ```

2. 解锁加锁的资源

   ```go
   mu.Lock()
   defer mu.Unlock()
   ```

3. 文件结尾的时候打印最终报告

   ```go
   defer fmt.Println()
   ```

4. 结束的时候关闭文件

   ```go
   package main
   
   import "fmt"
   
   func main() { // DBO 是 数据库所有者  (DataBase Owner)。是数据库默认的一个角色
   	doDBOperations()
   }
   
   func connectToDB() { //DB是数据库的意思，这不是一个确切的内置函数，而是一个为方便讲解的使用驼峰缩写法的自造函数
   	fmt.Println("ok, connected to db")
   }
   
   func disconnectFromDB() {
   	fmt.Println("ok, disconnected from db")
   }
   
   func doDBOperations() {
   	connectToDB()
   	fmt.Println("Defering the database disconnect.")
   	defer disconnectFromDB() //function called here with defer
   	fmt.Println("Doing some DB operations ...")
   	fmt.Println("Oops! some crash or network error ...")
   	fmt.Println("Returning from function here!")
   	return
   	//terminate the program
   	// deferred function executed here just before actually returning, even if
   	// there is a return or abnormal termination before
   	//终止程序
   	//在实际返回之前执行的延迟函数，即使
   	//之前有一个返回或异常终止
   }
   
   ```



###  defer实现代码追踪。



```go
package main

import "fmt"

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func a() {
	trace("a")//因为是string格式的输入值，所以要加""
	defer untrace("a")
	fmt.Println("in a")
}
//当然编译型语言在同一文件内不存在顺序。
func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a() // 
}

func main() {
    b() //此处是一个嵌套结构，意在b()中执行a()，因为defer会在b()完全执行之后才执行，所以defer untrace("b")是最后返回的
}
```

输出：

```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

### defer记录函数参数与返回值



```go
package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF //EOF是End-Of-File
}

func main() {
	func1("Go")
}

```



# Go的异常处理 defer, panic, recover

学习自：

[Go的异常处理 defer, panic, recover - 蝈蝈俊 - 博客园 (cnblogs.com)](https://www.cnblogs.com/ghj1976/archive/2013/02/11/2910114.html#:~:text=panic 是用来表示非常严重的不可恢复的错误的。 在Go语言中这是一个内置函数，接收一个interface,{}类型的值（也就是任何值了）作为参数。 panic的作用就像我们平常接触的异常。 不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。 所以，Go语言中的异常，那真的是异常了。)



Go语言追求简洁优雅，所以，在极个别的情况下，也就是说，遇到真正的异常的情况下（比如除数为0了）。才使用Go中引入的Exception处理：`defer`, `panic`, `recover`。



**这几个异常的使用场景可以这么简单描述**：Go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。

panic是恐慌的意思，在GO中用来表示nil的等级。

```go
package main
 
import "fmt"
 
func main(){
    defer func(){ // 必须要先声明defer，否则不能捕获到panic异常
        fmt.Println("c")
        if err:=recover();err!=nil{
            fmt.Println(err) // 这里的err其实就是panic传入的内容，55
        }
        fmt.Println("d")
    }()
    f()
}
 
func f(){
    fmt.Println("a")
    panic(55)
    fmt.Println("b")
    fmt.Println("f")
}
```





# panic

panic 英文原意：n. 恐慌，惊慌；大恐慌  adj. 恐慌的；没有理由的  vt. 使恐慌  vi. 十分惊慌

panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。

panic的作用就像我们平常接触的异常。

不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。

所以，Go语言中的异常，那真的是异常了。

你可以试试，调用panic看看，程序立马挂掉，然后Go运行时会打印出调用栈。
但是，关键的一点是，即使**函数执行的时候panic了**，函数不往下走了，运行时并不是立刻向上传递panic，**而是去运行defer**，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。

panic就是这么简单。抛出个真正意义上的异常。

 

# recover

recover 英文原意： vt. 恢复；弥补；重新获得  vi. 恢复；胜诉；重新得球  n. 还原至预备姿势

**上面说到，panic的函数并不会立刻返回，而是先defer，再返回。**

这时候（defer的时候），如果有办法将panic捕获到，并阻止panic传递，那就异常的处理机制就完善了。

Go语言提供了recover内置函数，前面提到，一旦panic，逻辑就会走到defer那，那我们就在defer那等着，调用recover函数将会捕获到当前的panic（如果有的话），被捕获到的panic就不会向上传递了，于是，世界恢复了和平。你可以干你想干的事情了。

不过要注意的是，recover之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。
