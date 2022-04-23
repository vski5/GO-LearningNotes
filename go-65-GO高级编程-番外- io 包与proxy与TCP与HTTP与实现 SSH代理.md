[神奇的 Golang-IO 包 - 熊喵君的博客 | PANDAYCHEN](https://pandaychen.github.io/2020/01/01/MAGIC-GO-IO-PACKAGE/)

io.copy传输，从A到B，缓冲区定长，所以对于不稳定的网络很友好。

`ioutil.ReadAll()`
`io.Copy()`
`ioutil.ReadAll()`



# io.Copy 与 Tcp 的开发结合

做过服务端开发的同学一定都写过代理 Proxy，代理的本质，是转发两个相同方向路径（但方向相反）上的 stream（数据流）。

例如，一个 的代理模式，B 作为代理，需要完成下面两件事情：`A-->B-->C`

1.  读取从 的数据，转发到 `A--->B``B--->C`
2.  读取从 的数据，转发到 `C--->B``B--->A`

在 golang 中，只需要 就能轻而易举的完成上面的事情，其 [实现代码](https://golang.org/src/io/io.go?s=12796:12856#L353) 如下所示：`io.Copy()`

```go
func copyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error) {
	// If the reader has a WriteTo method, use it to do the copy.
	// Avoids an allocation and a copy.
	if wt, ok := src.(WriterTo); ok {
		return wt.WriteTo(dst)
	}
	// Similarly, if the writer has a ReadFrom method, use it to do the copy.
	if rt, ok := dst.(ReaderFrom); ok {
		return rt.ReadFrom(src)
	}
	if buf == nil {
		size := 32 * 1024
		if l, ok := src.(*LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	for {
		nr, er := src.Read(buf)
		if nr > 0 {
			nw, ew := dst.Write(buf[0:nr])
			if nw > 0 {
				written += int64(nw)
			}
			if ew != nil {
				err = ew
				break
			}
			if nr != nw {
				err = ErrShortWrite
				break
			}
		}
		if er != nil {
			if er != EOF {
				err = er
			}
			break
		}
	}
	return written, err
}
```

#  io.Copy 与 http 开发相关
做出来的小项目在：[vski5/solid-rain：下载软件的迭代测试 (github.com)](https://github.com/vski5/solid-rain)
## 下载文件到本地：方法一
`http.Get()`
`ioutil.WriteFile()`
将下载内容直接写到文件中,如果是大文件，会出现内存不足的问题，因为它是需要先把请求内容全部读取到内存中，然后再写入到文件中的。
```go
func DownloadFile() error {
    url :="http://xxx/somebigfile"
    resp ,err := http.Get(url)
    if err != nil {
        fmt.Fprint(os.Stderr ,"get url error" , err)
    }

    defer resp.Body.Close()

    data ,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

	return ioutil.WriteFile("/tmp/xxx_file", data, 0755)
}
```

##：下载文件到本地：方法2，使用 `io.Copy()`

```go
func DownloadFile(url1, file2 string) {
	//url1 :="http://xxx/somebigfile"
	resp, err := http.Get(url1)
	if err != nil {
		fmt.Fprint(os.Stderr, "get url error", err)
	}
	defer resp.Body.Close()

	outfile, err2 := os.OpenFile(file2, os.O_RDWR|os.O_CREATE, os.ModePerm)//创造文件保存的位置，下载下来的文件会下到这个位置这个名字。
	if err2 != nil {
		fmt.Println(err2)
	}
	// 很重要：初始化一个 io.Writer类型的句柄，是转换类型也是在创造储存空间。
	wt := bufio.NewWriter(outfile)

	defer outfile.Close()

	jieshoudizhi, err := io.Copy(wt, resp.Body)
	if err != nil {
		panic(err)
	}
	//缓冲写入wt，目标文件的句柄，然后传递到正式的文件那里。
	wt.Flush()
	//防止jieshoudizhi出现没被利用的报错。
	fmt.Fprintln(os.Stdout, "no err", jieshoudizhi)

}
```

# io.copy与SSH代理。
需要连续用两个`io.Copy()`。
看上去像这样：`io.Copy()``io.Copy()`
