package main

import (
	"fmt"
	"log"
	"net/http"
)

// 处默认路由的页面 !!
func handler(w http.ResponseWriter, r *http.Request) { // r *http.Request 表示读取客户端的请求
	fmt.Fprintf(w, "Hi, 这里是 %s!", r.URL.Path[1:]) // 🌟 r.URL.Path[1:] 表示请求的 url 路径, 通过切片[1:]获取路径 / 之后的所有内容 => 比如获取 http://localhost:8080/gopher 的 gopher 🌟  fmt.Fprint 表示往 http.ResponseWriter 内写内容
}

func main() {
	// 指定路由
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
