package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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
