
## 通道

不仅可以使用原子函数和互斥锁来保证对共享资源的安全访问以及消除竞争状态，

还可以使用通道，通过发送和接收需要共享的资源，在 goroutine 之间做同步

当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了 确保同步交换数据的机制。

声明通道时，需要指定将要被共享的数据的类型。

可以通过通道共享
1. 内置类型、
2. 命名类型、
3. 结构类型、
4. 引用类型的值或者指针

# 通道
_通道(channels)_ 是连接多个协程的管道。 你可以从一个协程将值发送到通道，然后在另一个协程中接收。

默认发送和接收操作是阻塞的，直到发送方和接收方都就绪。 这个特性允许我们，不使用任何其它的同步操作， 就可以在程序结尾处等待消息 `"ping"`。

```go
package main

import "fmt"

func main() {
	//使用 make(chan val-type) 创建一个新的通道。 通道类型就是他们需要传递值的类型。
	messages := make(chan string)
	//使用 channel <- 语法 发送 一个新的值到通道中。
	//这里我们在一个新的协程中,用匿名函数发送 "ping" 到上面创建的 messages 通道中。
	go func() { messages <- "ping" }()
	//使用 <-channel 语法从通道中 接收 一个值。 这里我们会收到在上面发送的 "ping" 消息并将其打印出来。
	msg := <-messages
	fmt.Println(msg)
}

```

# 通道缓冲
默认情况下，通道是 _无缓冲_ 的，这意味着只有对应的接收（`<- chan`） 通道准备好接收时，才允许进行发送（`chan <-`）。 _有缓冲通道_ 允许在没有对应接收者的情况下，缓存一定数量的值。

```go
package main

import "fmt"

func main() {
	//创造有缓冲的通道，此处创造了能缓存两个string的通道
	message := make(chan string, 2)
	//为通道赋值
	message <- "test1"
	message <- "test2"
	message <- "test3" //即使这里赋予了第三个值，但只会输出两个值，这个通道只能缓冲前两个，最后一个不会顶掉第一个
	fmt.Println(<-message) //注意这里的通道表现方式，带上箭头才是完整的通道表现形式。
	fmt.Println(<-message)
	fmt.Println(<-message)
	//赋予两个值的时候，只输出一个的时候，会输出通道缓存的地址。例如0xc000050060

}

```

# 通道同步
我们可以使用通道来同步协程(goroutine)之间的执行状态。 

通道是goroutine之间通信的方式。

这儿有一个例子，使用阻塞接收的方式，实现了等待另一个协程完成。

如果需要等待多个协程,sync.WaitGroup是一个更好的选择。

```go
package main

import (
	"fmt"
	"time"
)

func worker(ok chan bool) {
	fmt.Println("STAT")
	time.Sleep(time.Second)
	fmt.Println("END")
	ok <- true
}
func main() {
	//为参数赋值
	ok := make(chan bool, 1)
	//运行一个协程，并给予用于通知的通道
	go worker(ok)
	//程序将一直阻塞，直至收到 worker 使用通道发送的通知。
	//也就是一直等到ok被传输结束。
	<-ok //这里居然要和自定义的worker()里的范例输入值一模一样才行，这个ok是阻塞“更底层”的worker()
}

```


# 通道方向
通道可以被规定是否为“只读（仅用于接收数据）”，“只写（只能发送数据）”。


```go
package main

import "fmt"

//test1函数定义了一个只能发送数据的（只写）通道
//从nsg到只写（只能发送）
func test1(write chan string, msg string) {
	write <- msg
}

//test2函数接收两个通道，read 仅用于接收数据（只读），write仅用于发送数据（只写）。
//从只写到msg，再从msg到只读（只能接收）
func test2(write chan string, read chan string) {
	msg := <-write
	read <- msg //这里的msg中转作用
}
func main() {
	t1 := make(chan string, 1)
	t2 := make(chan string, 1)
	test1(t1, "message")
	test2(t1, t2)
	fmt.Println(<-t2) //输出通道t2内的值,不写<- 就是输出地址。
}

```



# 通道选择器select,可用于阻塞和选择
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "1 is ok"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "2 is ok"
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

}

```
注意，程序总共仅运行了两秒左右。因为 1 秒 和 2 秒的 `Sleeps` 是并发执行的，

当select是用于阻塞，只要满足select中case里面任意条件就行的时候，冒号后面的执行语句可以写在一起。

但我忘记怎么写了，参考方应航的教程。
![[select.png]]
这是一个订阅是否有系统中断的select，中断后的执行是一样的，所以执行的操作写在最后面，相当于，在接收到想接收的条件后，只是简单的结束阻塞，剩下的操作交给其他的代码处理。
```go
select{
case<-条件:
case<-条件2:
}
结束阻塞后需要执行的语句。
```


# 基于channel和select的超时处理

这里是使用 `select` 实现一个超时操作。 `res := <- c1` 等待结果，`<-time.After` 等待超时（1秒钟）以后发送的值。 由于 `select` 默认处理第一个已准备好的接收操作， 因此如果操作耗时超过了允许的 1 秒的话，将会执行超时 case。
```go
select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }
```



```go
package main

import (
	"fmt"
	"time"
)

func main() { //程序运行时间大概再7S左右
	//超时的情况
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "1 is ok"
	}()
	select {
	case t1 := <-c1:
		fmt.Println(t1)
	case t1 := <-time.After(2 * time.Second):
		fmt.Println("time out", t1)
	}
	//准时的情况下
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "2 is ok"
	}()
	select {
	case t2 := <-c2:
		fmt.Println(t2)
	case t2 := <-time.After(10 * time.Second):
		fmt.Println("time out", t2)
	}

}

```
上面的代码，select在中间阻断着，所以两个func不是同时发生的，先跑了2S的等待，再跑5S的等待。

把select都放到最后，整个程序就只要大概7S，因为两个go func（也就是协程goroutine）是几乎同时就行的，

```go

package main

import (
	"fmt"
	"time"
)

func main() {
	//超时的情况
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "1 is ok"
	}()

	//准时的情况下
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "2 is ok"
	}()
	//把select都放到最后，不在前面阻塞

	select {
	case t1 := <-c1:
		fmt.Println(t1)
	case t1 := <-time.After(2 * time.Second):
		fmt.Println("time out", t1)
	}

	select {
	case t2 := <-c2:
		fmt.Println(t2)
	case t2 := <-time.After(10 * time.Second):
		fmt.Println("time out", t2)
	}

}

```

输出
>time out 2022-05-11 21:53:43.5634268 +0800 CST m=+2.008467101
2 is ok

改变select的顺序，前面的select会堵住后面的select，导致时间达到判断出1 is ok的要求。
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	//超时的情况
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "1 is ok"
	}()

	//准时的情况下
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(5 * time.Second)
		c2 <- "2 is ok"
	}()
	//把select都放到最后，不在前面阻塞
	select {
	case t2 := <-c2:
		fmt.Println(t2)
	case t2 := <-time.After(10 * time.Second):
		fmt.Println("time out", t2)
	}

	select {
	case t1 := <-c1:
		fmt.Println(t1)
	case t1 := <-time.After(2 * time.Second):
		fmt.Println("time out", t1)
	}

}

```

输出：
>2 is ok
1 is ok

# 非阻塞通道操作
常规的通过通道发送和接收数据是阻塞的。 然而，我们可以使用带一个 `default` 子句的 `select` 来实现 _非阻塞_ 的发送、接收，甚至是非阻塞的多路 `select`。

default是指在其他case都不复合的时候，直接执行default后面的代码。

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	/*
	   如果在 messages 中存在，然后 select 将这个值带入 <-messages case 中。
	   否则，就直接到 default 分支中。
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	/*
	   sg 不能被发送到 message 通道，因为这是 个无缓冲区通道，并且也没有接收者，因此， default 会执行。
	*/
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	/*
	   可以在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	   这里我们试图在 messages 和 signals 上同时使用非阻塞的接收操作。
	*/
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

```

# 通道的关闭
_关闭_ 一个通道意味着不能再向这个通道发送值了。

该特性可以向通道的接收方传达工作已经完成的信息。

使用` j, more := <- jobs` 配合`for`循环的从 通道 `jobs` 接收数据。

第二个值为ture，但当通道被历遍完全之后，第二个值就会改为false，这可作为判断，活用于if等条件语句后面，使通道关闭后的下一步任务得到启动。

`<-done` //阻塞，直到done这个通道接收到消息。（这运用了通道同步的思想）

```go
package main

import "fmt"

/*
使用一个 jobs 通道，将工作内容，从 main() 协程传递到一个工作协程中。
当我们没有更多的任务传递给工作协程时，我们将 close 这个 jobs 通道。
*/
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() { //这是工作协程
		for {
			j, more := <-jobs //使用 j, more := <- jobs 循环的从 jobs 接收数据。
			if more {
				/*根据接收的第二个值，如果 jobs 已经关闭了，
				并且通道中所有的值都已经接收完毕，那么 more 的值将是 false。
				当我们完成所有的任务时，会使用这个特性通过 done 通道通知 main 协程。*/
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()
	//使用 jobs 发送 3 个任务到工作协程中，然后关闭 jobs
	for j := 1; j <= 3; j++ {
		jobs <- j //将j值传递到通道jobs中
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done //阻塞，直到done这个通道接收到消息。
}

```

>D:\GOtext\three>go run .
sent job 1
sent job 2       
sent job 3       
sent all jobs    
received job 1   
received job 2   
received job 3   
received all jobs

jobs虽然能容纳5个缓存，但是只传了3个，所以只能历遍三个。

# 通道遍历
```go
package main

import "fmt"

func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    for elem := range queue {
        fmt.Println(elem)
    }
}
```

`range` 迭代从 `queue` 中得到每个值。 因为我们在前面 `close` 了这个通道，所以，这个迭代会在接收完 2 个值之后结束。

一个非空的通道也是可以关闭的， 并且，通道中剩下的值仍然可以被接收到。
