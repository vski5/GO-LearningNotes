# go-39-093-锁和 sync 包

# 上锁

上锁：变量被一个线程改变时（临界区）就是上锁，直到这个线程执行完成并解锁后，其他线程才能访问这个变量。

上锁用来解决：复杂程序中需要多线程执行不同应用实现程序的并发，不同线程使用同一个变量的时候，会导致无法预知变量被谁先更改，这被称为资源竞争，解决方法是上锁。

map类型不存在锁的机制，这提高了map类型的性能，也导致map类型是非线程安全的，不可并行访问map类型。

# 上锁就是用sync包中的Mutex

sync来自synchronized，意为同步的。这意味着线程将有序的对同一变量进行访问。

Mutex意为互斥（体）

`sync.Mutex` 是一个互斥锁，它的作用是守护在临界区入口来确保同一时间只能有一个线程进入临界区。

一个例子：

假设此中的info是一个需要上锁的放在共享内存中的变量

```go
import  "sync"

type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.: Str string
}
```

如果一个函数想要改变这个变量可以这样写:

```go
func Update(info *Info) {
	info.mu.Lock()
    // critical section:
    info.Str = // new value
    // end critical section
    info.mu.Unlock()
}
```

还有一个很有用的例子是通过 Mutex 来实现一个可以上锁的共享缓冲器:

```go
type SyncedBuffer struct {
	lock 	sync.Mutex
	buffer  bytes.Buffer
}
```

在 sync 包中还有一个 `RWMutex` 锁：他能通过 `RLock()` 来允许同一时间多个线程对变量进行读操作，但是只能一个线程进行写操作。如果使用 `Lock()` 将和普通的 `Mutex` 作用相同。包中还有一个方便的 `Once` 类型变量的方法 `once.Do(call)`，这个方法确保被调用函数只能被调用一次。

相对简单的情况下，通过使用 sync 包可以解决同一时间只能一个线程访问变量或 map 类型数据的问题。如果这种方式导致程序明显变慢或者引起其他问题，我们要重新思考来通过 goroutines 和 channels 来解决问题，这是在 Go 语言中所提倡用来实现并发的技术。我们将在第 14 章对其深入了解，并在第 14.7 节中对这两种方式进行比较。

# 来自go by example 的例子：

```go
package main

import (
	"fmt"
	"sync"
)

/*
Container 中定义了 counters 的 map ，由于我们希望从多个 goroutine 同时更新它，
因此我们添加了一个 互斥锁Mutex 来同步访问。
请注意不能复制互斥锁，
如果需要传递这个 struct，应使用指针完成。
*/
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

/*
在访问 counters 之前锁定互斥锁；
使用 [defer]（defer） 在函数结束时解锁。
*/
func (c *Container) inc(name string) { //如果需要传递这个 struct，应使用指针完成。

	c.mu.Lock()         //上锁
	defer c.mu.Unlock() //解锁
	//defer是推迟语句，会在整个执行完后执行。
	c.counters[name]++
}

/*
请注意，互斥量的零值是可用的，因此这里不需要初始化。
*/
func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	//这个函数在循环中递增对 name 的计数
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}
	/*
	   同时运行多个 goroutines; 请注意，它们都访问相同的 Container，其中两个访问相同的计数器。
	*/
	wg.Add(3)
	go doIncrement("a", 100)
	go doIncrement("a", 10000)
	go doIncrement("b", 2000)

	wg.Wait() //等待上面的 goroutines 都执行结束
	fmt.Println(c.counters)
}

```

输出：

```
map[a:10100 b:2000]
```

