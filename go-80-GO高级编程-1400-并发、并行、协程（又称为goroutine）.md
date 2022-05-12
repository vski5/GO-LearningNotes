

**不要通过共享内存来通信，而通过通信来共享内存。**
## 线程、进程
线程（thread）和进程（process）

一个线程是一个执行空间，这个空间会被操作系统调度来运行 函数中所写的代码。

进程：分配资源的独立执行体。包含了应用程序在运行中需要用到和维护的各种资源的容器

每个进程至少包含一个线程，每个进程的初始线程被称作主线程。因为执行 这个线程的空间是应用程序的本身的空间，所以当主线程终止时，应用程序也会终止。


一个运行的应用程序的进程和线程的简要描绘

![[一个运行的应用程序的进程和线程的简要描绘.png]]



## 并发、并行、协程


操作系统会在物理处理器上调度线程来运行，而 Go 语言的运行时会在逻辑处理器上调度 goroutine来运行。

每个逻辑处理器都分别绑定到单个操作系统线程。

默认给 一整个应用程序只分配一个逻辑处理器。

这些逻辑处理器会用于执行所有被创建的goroutine。并发调度无数个goroutine

Go 调度器如何管理 goroutine
![[Go 调度器如何管理 goroutine.png]]

并发程序可以在一个处理器或者内核上使用多个线程来执行任务，

但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。

并行是一种通过使用多处理器以提高速度的能力。

所以并发程序可以是并行的，也可以不是。


![[并发和并行的区别.png]]



## goroutine

![[goroutine 在逻辑处理器的线程上进行交换.png]]

在第 1 步，调度器开始运行 goroutine A，而 goroutine B 在运行队列里等待调度。

在第 2 步，调度器交换了 goroutine A 和 goroutine B。 由于 goroutine A 并没有完成工作，因此被放回到运行队列。

在第 3 步，goroutine B 完成 了它的工作并被系统销毁。这也让 goroutine A 继续之前的工作。

一个例子：
创建了两个 goroutine，分别打印 1~5000 内的素数。查找并显示素数 会消耗不少时间，这会让调度器有机会在第一个 goroutine 找到所有素数之前，切换该 goroutine 的时间片。
goroutine B 先显示素数。一旦 goroutine B 打印到素数 4591，调度器就会将正运行的 goroutine 切换为 goroutine A。之后 goroutine A 在线程上执行了一段时间，再次切换为 goroutine B。这次 goroutine B 完成了所有的工作。一旦 goroutine B 返回，就会看到线程再次切换到 goroutine A 并 完成所有的工作。每次运行这个程序，调度器切换的时间点都会稍微有些不同。
```go
//这个示例程序展示 goroutine 调度器是如何在单个线程上
//切分时间片.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// main is the entry point for all Go programs.
func main() {
	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(1)

	// Add a count of two, one for each goroutine.
	wg.Add(2)

	// Create two goroutines.
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示前5000个数字的质数
func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

```

## 如何修改逻辑处理器的数量

	// Allocate 1 logical processors for the scheduler to use.
	runtime.GOMAXPROCS(1)

    // 给每个可用的核心分配一个逻辑处理器 
    runtime.GOMAXPROCS(runtime.NumCPU())

包 runtime 提供了修改 Go 语言运行时配置参数的能力。


# gobyexample-cn.github.io/goroutines

```go
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct") //同步的调用此函数

	go f("goroutine") //把f()加入并发队列，会并发执行

	go func(msg string) { //并发执行匿名函数
		fmt.Println(msg)
	}("going") //匿名函数的参数

	time.Sleep(time.Second) //睡一秒，模拟sync.WaitGroup的效果
	//现在两个协程在独立的协程中 异步地 运行， 然后等待两个协程完成
	fmt.Println("done")
}

```


运行这个程序时，首先会看到阻塞式调用的输出，(同步调用的函数，不在并发队列的函数先一口气输出完)。
然后是两个协程的交替输出。这种交替的情况表示 Go runtime 是以并发的方式运行协程的。
输出：
>
direct : 0
direct : 1
direct : 2
going
goroutine : 0
goroutine : 1
goroutine : 2
done