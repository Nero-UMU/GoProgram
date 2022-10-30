package main

import (
	"fmt"
	"net/http"
)

func msgHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	hello := r.URL.Query().Get("hello")
	if hello == "world" {
		fmt.Fprintf(w, "fuck me")
	} else {
		fmt.Fprintf(w, "hello world")
	}
}

func main() {
	http.HandleFunc("/", msgHandler)

	// 生成一个文件服务器,根目录是 public 目录
	fs := http.FileServer(http.Dir("/public"))
	// 生成一个处理程序,第二个参数是一个返回 handle 接口的函数
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.ListenAndServe(":8080", nil)
}
