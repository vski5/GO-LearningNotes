# go-130-GO-WEB编程-1

 [webv1.go](./GO-WEB编程/webv1.go)
```go

package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

```

-   首先调用`Http.HandleFunc`
    
    按顺序做了几件事：
    
    1 调用了DefaultServeMux的HandleFunc
    
    2 调用了DefaultServeMux的Handle
    
    3 往DefaultServeMux的map[string]muxEntry中增加对应的handler和路由规则
    
-   其次调用`http.ListenAndServe(":9090", nil)`
    
    按顺序做了几件事情：
    
    1 实例化Server
    
    2 调用Server的ListenAndServe()
    
    3 调用net.Listen("tcp", addr)监听端口
    
    4 启动一个for循环，在循环体中Accept请求
    
    5 对每个请求实例化一个Conn，并且开启一个goroutine为这个请求进行服务go c.serve()
    
    6 读取每个请求的内容w, err := c.readRequest()
    
    7 判断handler是否为空，如果没有设置handler（这个例子就没有设置handler），handler就设置为DefaultServeMux
    
    8 调用handler的ServeHttp
    
    9 在这个例子中，下面就进入到DefaultServeMux.ServeHttp
    
    10 根据request选择handler，并且进入到这个handler的ServeHTTP
    
      mux.handler(r).ServeHTTP(w, r)
    
    11 选择handler：
    
    A 判断是否有路由能满足这个request（循环遍历ServeMux的muxEntry）
    
    B 如果有路由满足，调用这个路由handler的ServeHTTP
    
    C 如果没有路由满足，调用NotFoundHandler的ServeHTTP
三步让web运行起来：
1.监听端口
```go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}

```
`ListenAndServe`调用了`net.Listen("tcp", addr)`，也就是底层用TCP协议搭建了一个服务，最后调用监控我们设置的端口。

2.接受客户端请求
Conn：用户的每次请求链接
这个函数里面起了一个`for{}`，首先通过`Listener`接收请求：`l.Accept()`，其次创建一个`Conn：c := srv.newConn(rw)`，最后单独开了一个`goroutine`，`go c.serve(connCtx)`把这个请求的数据当做参数扔给这个conn去服务

3.分配handler

conn首先会解析request: `w, err := c.readRequest(ctx)`, 然后获取相应的handler`ServeHTTP`去处理请求`serverHandler{c.server}.ServeHTTP(w, w.req)`

```go
func (sh serverHandler) ServeHTTP(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}
	if req.RequestURI == "*" && req.Method == "OPTIONS" {
		handler = globalOptionsHandler{}
	}
	handler.ServeHTTP(rw, req)
}
```

