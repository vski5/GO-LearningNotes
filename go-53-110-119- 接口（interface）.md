# 接口：简介 
**接口不就是对一些方法的约束集合吗**

**接口** ：通过它可以实现很多面向对象的特性。
接口提供了一种方式来 **说明** 对象的行为：如果谁能搞定这件事，它就可以用在这儿。

接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

通过如下格式定义接口：

```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```

上面的 `Namer` 是一个 **接口类型**。

（按照约定，只包含一个方法的）接口的名字由方法名加 `[e]r` 后缀组成，例如 `Printer`、`Reader`、`Writer`、`Logger`、`Converter` 等等。还有一些不常用的方式（当后缀 `er` 不合适时），比如 `Recoverable`，此时接口名以 `able` 结尾，或者以 `I` 开头（像 `.NET` 或 `Java` 中那样）。

Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。

## 标准库里有很好的例子
如 io 包里 实现的流式处理接口。

io 包提供了一组构造得非常好的接口和函数，来让代码轻松支持流式数据处理。

只要实现两个接口，就能利用整个 io 包背后的所有强大能力。 

# 标准库
## 这个程序实现了流行程序 curl(命令行工具)的功能
curl功能是：基于网络协议，对指定URL进行网络传输。
代码如下：
```go

// Sample program to show how to write a simple version of curl using

  

// the io.Reader and io.Writer interface support.

// 这个示例程序展示如何使用 io.Reader 和 io.Writer 接口

// 写一个简单版本的 curl 程序

package main

  

import (

 "fmt"

 "io"

 "net/http"

 "os"

)

  

// init is called before main.

  

func init() {

  

 if len(os.Args) != 2 { //Args保存命令行参数，从程序名开始

 /*

 var Args []string

 Args保管了命令行参数，第一个是程序名。

 */

 fmt.Println("Usage: ./example2 <url>")

  

 os.Exit(-1)

 //os.Exit可以让当前程序使用给定的状态代码退出。

 //按照惯例，代码0表示成功，非0表示错误。程序立即终止;不运行延迟函数。

  

 }

  

}

  

// main is the entry point for the application.

  

func main() {

  

 // Get a response from the web server.

 // 从 Web 服务器得到响应

 /*

 在第下一行，调用了 http 包的 Get 函数。

 在与服务器成功通信后，http.Get 函数会返回一个 http.Response 类型的指针。

 http.Response 类型包含一个名为 Body 的字段，这个字段是一个 io.ReadCloser 接口类型的值

 */

 r, err := http.Get(os.Args[1])

 //r是http.Response 类型的指针

 if err != nil {

 fmt.Println(err)

 return

 }

  

 // Copies from the Body to Stdout.

 // 从 Body 复制到 Stdout

 /*

 Body 字段作为第二个参数传给 io.Copy 函数。

 io.Copy 函数的第二个参数，接受一个 io.Reader 接口类型的值，这个值表示数据流入的源。

 Body 字段实现了 io.Reader接口，因此我们可以将 Body 字段传入 io.Copy，使用 Web 服务器的返回内容作为源。

 io.Copy 的第一个参数是复制到的目标，这个参数必须是一个实现了 io.Writer 接口的值。

 对于这个目标，我们传入了 os 包里的一个特殊值 Stdout。

 这个接口值表示标准输出设备，并且已经实现了 io.Writer 接口。

 当我们将 Body 和 Stdout 这两个值传给 io.Copy 函数后，

 这个函数会把服务器的数据分成小段，源源不断地传给终端窗口，直到最后一个片段读取并写入终端，io.Copy 函数才返回。

 */

 io.Copy(os.Stdout, r.Body)

 //Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。

  

 if err := r.Body.Close(); err != nil {

 fmt.Println(err)

 }

  

}

```


### `func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)`

简单地讲，是把`src io.Reader`的值传入复制给`dst io.Writer`。
然后当我们将`dst`和`src `这两个值传给` io.Copy `函数后， 这个函数会把服务器的数据分成小段，源源不断地传给终端窗口，直到最后一个片段读取并写入 终端，`io.Copy `函数才返回。

1. 第一个参数
	
	**io.Copy 的第一个参数`dst io.Writer`是复制到的目标，**
	这个参数必须是一个实现了 io.Writer 接口的值，
	在下面的例子里，我们传入了 os 包里的一个特殊值 Stdout，（第一个参数被设为了 Stdout）。
	**Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。**
	这个接口值`Stdout`表示标准输出设备， 并且已经实现了` io.Writer `接口。
	第一个参数是实现了` io.Writer `接口的，有三种，Stdin, Stdout,  Stderr 。
	
2. 第二个参数
	
	io.Copy 函数的第二个参数， 接受一个 io.Reader 接口类型的值，
	**这个值表示数据流入的源。**
	`src io.Reader`实现了 io.Reader 接口，因此我们可以将`src io.Reader`字段传入 io.Copy，使用 Web 服务器的返回内容作为源。


在上面的例子中，当我们将 Body 和 Stdout 这两个值传给 io.Copy 函数后， 这个函数会把服务器的数据分成小段，源源不断地传给终端窗口，直到最后一个片段读取并写入 终端，io.Copy 函数才返回。

## io.Copy 函数可以以这种工作流的方式处理很多标准库里已有的类型
具体例子如下所示:
```go
// Sample program to show how a bytes.Buffer can also be used

// with the io.Copy function.

package main

  

import (

 "bytes"

 "fmt"

 "io"

 "os"

)

  

// main is the entry point for the application.

func main() {

 var b bytes.Buffer

  

 // Write a string to the buffer.

 b.Write([]byte("Hello"))

  

 // Use Fprintf to concatenate a string to the Buffer.

 fmt.Fprintf(&b, "World!")

  

 // Write the content of the Buffer to stdout.

 io.Copy(os.Stdout, &b)

 //func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)

 //使用 io.Copy 函数，将字符写到终端窗口。

}

```


# 实现
接口也是一种类型。

接口是方法的集合，实现接口，就需要实现接口内的所有方法。

如果：你定义的类型实现了某个接口里声明的所有方法，那这个类型才能赋值给这个接口。赋值之后，你定义的类型就会存入接口，成为接口类型的值之一。

实体类型：用户定义的类型，离开内部储存的实体类型的话，接口值不会有任何具体的行为。

恰当的实体类型可以实现任何恰当的接口，对接口值方法的调用会是一种多对一的情况。

## 用户定义的类型的值或者指针要满足接口的实现，需要遵守一些 规则。
一个例子
```go
var n notifier //类型为 notifier 名为 n 的接口变量
n = user("bill") //即使是把指针赋值给接口 n = &user("bill") 内部布局不变。
```
在 一个名为 user 的 类型值 赋值后 接口变量的值 的 内部布局如下所示
接口值是一个 两个字长度 的数据结构：
1. 指向内部表(iTable)的指针，指针揭示地址。iTable包含user的类型 和 方法集
2. 指向user值的指针，指针揭示地址。user是储存的值。


# 方法集
先解释方法集：方法集定义了接口的接受规则，接口能传什么值，全靠方法集规定。

另外一个前情提要：
Go语言里定义的方法集的规则是：  
1. 从值的角度来看规则
	Values------------Methods Receivers	
	T----------------------(t T)	
	\*T------------ (t T) and (t \*T)	

T类型的值的方法集只包含值接收者声明的方法。
而指向T类型的指针的方法集既包含值接收者声明的方法，也包含指针接收者声明的方法。

2. 从接收者的角度来看规则

	Values------------Methods Receivers
	(t T)--------------------T and \*T
	(t \*T) ------------------------   \*T

使用指针类型的接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。
如果使用值类型的接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。


```go
// Sample program to show how to use an interface in Go.

// Sample program to show how to use an interface in Go.

package main

  

import (

 "fmt"

)

  

// notifier is an interface that defined notification

// type behavior.

type notifier interface {

 notify()

}

  

// user defines a user in the program.

type user struct {

 name string

 email string

}

  

// notify implements a method with a pointer receiver.

func (u *user) notify() {

 fmt.Printf("Sending user email to %s<%s>\n",

 u.name,

 u.email)

}

  

// main is the entry point for the application.

func main() {

 // Create a value of type User and send a notification.

 u := user{"Bill", "bill@email.com"}

 //上文声明了，u是*user，一个指针。

 sendNotification(&u)

 /*不能使用u(类型为用户的变量)作为参数sendNotification的通知值

 user 类型的值并没有实现这个接口。

 上文声明了，u是*user，一个指针。

 需要改为 &u 即 u的内容。

 此处u是接收者

 使用指针类型的接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。

 如果使用值类型的接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。

 */

  

}

  

// sendNotification accepts values that implement the notifier

// interface and sends notifications.

func sendNotification(n notifier) {

 n.notify()

}

```



## 多态

```go

// Sample program to show how polymorphic behavior with interfaces.

package main

  

import (

 "fmt"

)

  

// notifier is an interface that defines notification

// type behavior.

type notifier interface {

 notify()

}

  

// user defines a user in the program.

type user struct {

 name string

 email string

}

  

// notify implements the notifier interface with a pointer receiver.

func (u *user) notify() {

 fmt.Printf("Sending user email to %s<%s>\n",

 u.name,

 u.email)

}

  

// admin defines a admin in the program.

type admin struct {

 name string

 email string

}

  

// notify implements the notifier interface with a pointer receiver.

func (a *admin) notify() {

 fmt.Printf("Sending admin email to %s<%s>\n",

 a.name,

 a.email)

}

  

// main is the entry point for the application.

func main() {

 // Create a user value and pass it to sendNotification.

 bill := user{"Bill", "bill@email.com"}

 sendNotification(&bill)

/*

我们两次声明了多态函数 sendNotification，

这个函数接受一个实现了notifier 接口的值作为参数。

既然任意一个实体类型都能实现该接口，那么这个函数可以针对任意实体类型的值来执行 notifier 方法。

因此，这个函数就能提供多态的行为

*/

 // Create an admin value and pass it to sendNotification.

 lisa := admin{"Lisa", "lisa@email.com"}

 sendNotification(&lisa)

}

  

// sendNotification accepts values that implement the notifier

// interface and sends notifications.

func sendNotification(n notifier) {

 n.notify()

}

```
