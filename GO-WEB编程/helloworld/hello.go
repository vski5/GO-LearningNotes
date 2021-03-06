package main

import (
	"fmt"
	"net/http"
)

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
