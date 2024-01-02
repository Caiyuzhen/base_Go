package main

import (
	"fmt"
	"log"
	"net/http"
    "io/ioutil"
)


// 处理 GET 请求
func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Get" {
		fmt.Fprintf(w, "👍 这是 GET 请求返回的数据")
	}
}


// 处理 POST 请求
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// 读取请求 body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("❌ 请求体读取失败", err)
		}

		// 读取成功, 输出 body 或做后续的处理
		fmt.Fprintf(w, "👍 成功读取到了请求体: %v", string(body))
	}
}


func main() {
	// 🔥 设置路由
	http.HandleFunc("/getLala", getHandler)
	http.HandleFunc("/postLala", postHandler)

	// 🔥 尝试在 8080 端口启动服务
	fmt.Println("尝试在8080端口启动服务器...")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("❌ 服务器启动失败:", err)
	} 
}