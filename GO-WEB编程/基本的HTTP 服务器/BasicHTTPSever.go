package main

import (
	"fmt"
	"net/http"
)

//处理动态请求
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {
	//处理动态请求
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//HandleFunc为 "DefaultServeMux "中给定的 "pattern"注册了处理函数。
	http.HandleFunc("/", RequestHandler)
	//处理静态资源
	fs := http.FileServer(http.Dir("static/"))
	//为/static/注册处理函数
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	//监听连接端口
	http.ListenAndServe("127.0.0.1:80", nil)
	select {}
}
