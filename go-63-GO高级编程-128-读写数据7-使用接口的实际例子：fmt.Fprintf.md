# 使用接口的实际例子：fmt.Fprintf

```go

package main

import (
	"bufio"
	"fmt" // interfaces being used in the GO-package fmt
	"os"
)

func main() {
	// unbuffered,无缓冲的
	fmt.Fprintf(os.Stdout, "%s\n", "hello world! - unbuffered")
	// buffered: os.Stdout 实现了 io.Writer类型
	// buffered缓冲(的过去分词)
	buf := bufio.NewWriter(os.Stdout) //用bufio.NewWriter()写到缓冲里面去。
	// and now so does buf.
	fmt.Fprintf(buf, "%s\n", "hello world! - buffered")
	buf.Flush() //Flush将所有缓冲数据写入底层io.Writer类型。必须写进去，不然报错。

}

```

输出：

```
hello world! - unbuffered
hello world! - buffered
```

下面是 `fmt.Fprintf()` 函数的实际签名

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

下面是io.Writer的定义
```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

 `fmt.Fprintf()` 函数,其不是写入一个文件，而是写入一个名为`io.Writer` 的接口类型的变量，标准输入输出错误都能实现这个接口类型。

fmt.Fprintf()指定的格式中第一个参数是实现了io.Writer接口的字符串，这个字符串可以是一个文件，管道，网络连接，通道等等。

io.Writer这个接口由一个Write方法组成，
bufio.Writer 实现了 Write 方法，所以可以作为Write方法的接收器。
```go
func (b *Writer) Write(p []byte) (nn int, err error)
```


它还有一个工厂函数：传给它一个 `io.Writer` 类型的参数，它会返回一个带缓冲的 `bufio.Writer` 类型的 `io.Writer`:

```go
func NewWriter(wr io.Writer) (b *Writer)
```

其适合任何形式的缓冲写入。

在缓冲写入的最后千万不要忘了使用 `Flush()`，否则最后的输出不会被写入。




在 15.2-15.8 章节，我们将使用 `fmt.Fprint` 函数向 `http.ResponseWriter` 写入，其同样实现了 io.Writer 接口。



**练习 12.7**：[remove_3till5char.go](app://obsidian.md/exercises/chapter_12/remove_3till5char.go)

下面的代码有一个输入文件 `goprogram`，然后以每一行为单位读取，从读取的当前行中截取第 3 到第 5 的字节写入另一个文件。然而当你运行这个程序，输出的文件却是个空文件。找出程序逻辑中的 bug，修正它并测试。



```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
)

func main() {
	inputFile, _ := os.Open("goprogram")
	outputFile, _ := os.OpenFile("goprogramT", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			return
		}
		outputString := string(inputString[2:5]) + "\r\n"
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Conversion done")
}
```



文件结构为
bufreadfile
 ├─goprogram.txt
 ├─bufread.go

所以把os.Open改成具体的文件（包括路径的那种）
`inputFile, _ := os.Open("bufreadfile/goprogram.txt")`

修改结果就是少了输出文件的句柄.Flush()，此处为outputWriter.Flush() ，这可以用来把缓冲区的内容完全写入文件。

```go
package main  
  
import (  
   "bufio"  
 "fmt" "os")  
  
func main() {  
   inputFile, _ := os.Open("bufreadfile/goprogram.txt")  
   outputFile, _ := os.OpenFile("goprogramT.txt", os.O_WRONLY|os.O_CREATE, 0666)  
  
   defer inputFile.Close()  
   defer outputFile.Close()  
   inputReader := bufio.NewReader(inputFile)  
   outputWriter := bufio.NewWriter(outputFile)  
   for {  
      inputString, readerError := inputReader.ReadString('\n')  
  
      outputWriter.Flush() //把缓冲区的内容完全写入文件  
  
 if readerError != nil { //一开始猜测是因为EOF这个检查错误的机制，就是拷贝不出最后一行的东西。换成nil检测机制一样有问题。  
 fmt.Println("EOF")  
         return  
 }  
      //outputString := string(inputString[:]) + "\r\n" //是用切片写入输出文件的问题，切片左开右闭，最后一项就读不出来。  
 outputString := string(inputString) + "\r\n" //直接读的话，就会进入循环，一行内容疯狂历遍。  
 //结论是直接用[[go-59-GO高级编程-123-读写数据3-文件拷贝]]里的io包怼上去。  
 _, err := outputWriter.WriteString(outputString)  
      if err != nil {  
         fmt.Println(err)  
         return  
 }  
  
   }  
   fmt.Println("Conversion done")  
}
```
结论是直接用[[go-59-GO高级编程-123-读写数据3-文件拷贝-io.copy粗读]]里的io包怼上去。  

```go
	n, e := io.Copy(file2, file1) //将file1拷贝到file2
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("拷贝成功。。。，拷贝字节数：", n)
	}
	
}
```


