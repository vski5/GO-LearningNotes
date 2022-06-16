# go-130-GO-WEB编程-4-Go搭建一个Web服务器
**本章是为了介绍如何-处理动态请求-接受链接**
**net/http包**：通过http包可以很方便的搭建起来一个可以运行的Web服务。同时使用这个包能很简单地对Web的路由，静态文件，模版，cookie等数据进行设置和操作。

## Request Handler
首先创造一个handler（处理函数/处理程序），怎样的处理函数呢，用于接收所有来自 浏览器或HTTP 客户端或 API 请求的 HTTP 连接。

```go
func (w http.ResponseWriter, r *http.Request)
```
该函数接收两个参数。
- `http.ResponseWriter` 写 text/html response 的地方，
- `http.Request` 包含了关于这个HTTP请求的所有信息，包括像URL或 header 

完善一下这个处理函数，
使得在默认的HTTP服务器上注册一个request handler请求处理程序：
`http.HandleFunc`会在下一节解释作用。
```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
})
```

## 倾听HTTP连接
```go
http.ListenAndServe(":80", nil)
```
启动Go的默认HTTP服务器并监听80端口的连接。
在80前加上127.0.0.1之类的本地地址，会因为没写阻塞直接寄。写了的话记得加上`select{}`

```go
func ListenAndServe(addr string, handler Handler) error
```
ListenAndServe监听TCP网络地址addr，然后用handler调用Serve来处理传入连接的请求,处理程序通常为nil。

接受的连接被配置为启用TCP keep-alives。

# 结果：
```go
package main  
  
import (  
   "fmt"  
 "net/http")  
  
func RequestHandler(w http.ResponseWriter, r *http.Request) {  
   fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)  
}  
  
func main() {  
   //func HandleFunc(pattern string, handler func(ResponseWriter, *Request))  
 //HandleFunc为 "DefaultServeMux "中给定的 "pattern"注册了处理函数。  
 http.HandleFunc("/", RequestHandler)  
   http.ListenAndServe("127.0.0.1:80", nil)  
   select {}  
}
```
