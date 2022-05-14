http 是比 tcp 更高层的协议，它描述了网页服务器如何与客户端浏览器进行通信。

Go 提供了 `net/http` 包.

回忆go-90：
1. 服务器的监听：`net.Listen("tcp", "localhost:50000")`


在 `net/http` 包中：
1. http.URL用于表示网页地址
2. 其中字符串属性 `Path` 用于保存 url 的路径
3. `http.Request` 描述了客户端请求，内含一个 `URL` 字段

一个例子，详解在后面：
```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
```

# HTTP 客户端
```go

```

# HTTP服务端


