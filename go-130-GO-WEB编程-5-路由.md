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

