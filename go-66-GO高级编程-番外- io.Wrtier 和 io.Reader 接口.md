一个程序的输出可以是另一个程序的输入，

这些程序使用 stdin 和 stdout 设备作为通道，在进程之间传递数据。

io包可以以流的方式高效处理数据，而不用考虑数据是什么，数据来自哪里，以及数据要发送到哪里的问题。

与 stdout 和 stdin 对应，这个包含有 io.Writer 和 io.Reader 两个接口。


# io.Writer，向数据流写入数据。

```go
type Writer interface { 
	Write(p []byte) (n int, err error) 
	}
```

从[]byte类型的p中向底层数据流写入数据。
返回写入的字节数n int 也就是 len(p)  。

如果写入的字节数小于len(p)，可能是任何导致写入提前结束的错误，就会返回err。

# io.Reader，从数据流中读取数据

```go
type Reader interface { 
	Read(p []byte) (n int, err error) 
	}
```

接受一个[]byte类型的p， Read 最多读入 len(p)字节，保存到 p。
允许出现读取的字节数小于 byte 切片的长度（也会占用所有p的空间存储临时数据），并且如果在读取时已 经读到数据但是数据不足以填满 byte 切片时，不应该等待新数据，而是要直接返回已读数据。

达到文件末尾（EOF）的情况：当读到最后一个字节 时，可以有两种选择：
1.  Read 返回最终读到的字节数，并且返回 EOF 作为错误值，
2. 返回最终读到的字节数，并返回 nil 作为错误值。下一次读取的时候， 由于没有更多的数据可供读取，需要返回 0 作为读到的字节数，以及 EOF 作为错误值

任何时候 Read 返回了读取的字节数，都应该优先 处理这些读取到的字节，再去检查 EOF 错误值或者其他错误值。

第四条约束建议 Read 方法的实现永远不要返回 0 个读取字节的同时返回 nil 作为错误值。

如果没有读到值，Read 应 该总是返回一个错误

返回两个值:
第一个值是读入的字节数，第二个值是 error 错误值。

# 总结 与 案例

 与 stdout 和 stdin 对应，这个包含有 io.Writer 和 io.Reader 两个接口。

从stdout进入 流 传出到 stdin 。（这两者为空间的实际内容）
从 io.Writer进入 流 传出到 io.Reader 。
如同隧道的两端。

 bytes.Buffer  是 一个实现了io.writer接口的struct。

## 一个例子

```go
// Sample program to show how different functions from the
// standard library use the io.Writer interface.
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main is the entry point for the application.
func main() {
	// 创建一个 Buffer 值，并将一个字符串写入 Buffer
	// 使用实现 io.Writer 的 Write 方法
	var b bytes.Buffer //一个实现了io.writer接口的struct。
	b.Write([]byte("Hello "))// Write 方法从[]byte("Hello ")中向底层数据流写入数据。此处的底层流是b bytes.Buffe（一个实现了io.writer接口的struct。）

	// Use Fprintf to concatenate a string to the Buffer.
	// Passing the address of a bytes.Buffer value for io.Writer.
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容输出到标准输出设备
	// 将 os.File 值的地址作为 io.Writer 类型值传入
	b.WriteTo(os.Stdout)
}

```
解析：
```go
// Write 将 p 的内容追加到缓冲区，如果需要，会增大缓冲区的空间。返回值 n 是 p 的长度，err 总是 nil。
//如果缓冲区变得太大，Write 会引起崩溃…
func (b *Buffer) Write(p []byte) (n int, err error) { 
	b.lastRead = opInvalid
	m := b.grow(len(p))
	return copy(b.buf[m:], p), nil
} 
```
由于实现了这个方法， 指向 Buffer 类型的指针b就满足了 io.Writer 接口，可以将指针作为第一个参数传入 Fprintf。

使用 Fprintf 函数，最终通过 Buffer 实现的 Write 方法， 将"World!"字符串追加到 Buffer 类型变量的内部缓冲区。（ bytes.Buffer  是 一个实现了io.writer接口的struct，Buffer类型变量有一个内部缓冲区。）


将整个 Buffer 类型变量 的内容写到 stdout。
```go
// 将 Buffer 的内容输出到标准输出设备
// 将 os.File 值的地址作为 io.Writer 类型值传入
b.WriteTo(os.Stdout)
```
使用 WriteTo 方法将 Buffer 类型的变量b的内容写到 stdout 设备。
这个方法接受一个实现了 io.Writer 接口的值。
在这个程序里，传入的值是 os 包的 Stdout 变量的值。

















































