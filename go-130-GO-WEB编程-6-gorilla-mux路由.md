# Routing (using gorilla/mux)
`gorilla/mux`可以实现复杂的请求路由。
本节用这个包来创建带有命名参数、GET/POST处理程序和域名限制的路由。
```sh
go get -u github.com/gorilla/mux
```

# 创建新的路由

首先创建一个新的请求路由器。该路由器是你的Web应用程序的主要路由器，以后将作为参数传递给服务器。它将接收所有的HTTP连接并将其传递给你将在其上注册的请求处理程序。
你可以像这样创建一个新的路由器。
```go
r := mux.NewRouter()
```

# 注册请求处理程序。
不是调用 `http.HandleFunc(...)` ，
而是像这样在你的路由器上调用：`HandleFuncr.HandleFunc(...)`

# URL参数
`gorilla/mux`可以实现从请求 URL 中提取段.
以URL`/books/go-programming-blueprint/page/10`为例
这个URL有两个dynamic segments(动态段/动态细分):
- /books/go-programming-blueprint。/books里的各种书名
- /page/10。 /page里的各种页面

为了让请求处理程序与上面提到的URL相匹配，你可以在你的URL模式中用占位符替换动态段，就像这样：
```go
r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    // get the book
    // navigate to the page
})
```

然后从这些片段中获取数据。
`gorilla/mux`提供了一个函数`mux.Vars(r)`，该函数将段`http.Request`作为参数，并返回一个段的map。

```go
func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    vars["title"] // the book title slug
    vars["page"] // the page
}
```


# 总结：
```go
package main  
  
import (  
   "fmt"  
 "net/http"  
 "github.com/gorilla/mux")  
  
func main() {  
   //建立新路由  
 r := mux.NewRouter()  
   //处理动态请求并处理静态资源。  
 //HandleFunc用URL路径的匹配器注册了一个新的路由。  
 r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {  
      vars := mux.Vars(r)    //mux中的特殊的map，Vars返回当前请求的路由变量(请求头*http.Request)，如果有的话。  
 title := vars["title"] //声明url中占位符对应的变量。  url是后面括号内的。
 page := vars["page"]  
      fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)  
   })  
  
   http.ListenAndServe(":80", r)  
}
```

# 特点：
## 方法:都作用在r.HandleFunc()上
### 将请求处理程序限制在特定的HTTP方法上。
```go
r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")
```
### 将请求处理程序限制在特定的 主机名 和 子域
```go
r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")
```
### 将请求处理程序限制为 http/https。
```go
r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")
```
### 路径前缀和子路由器
将请求处理程序限制在特定的路径前缀上
```go
bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
```

