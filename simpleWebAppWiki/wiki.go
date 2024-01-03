package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
)

// 定义 Pagr 的数据结构, 用于存储页面的标题和内容
type Page struct {
	Title string
	Body []byte // 为了方便存储, 不定义为 string, 而是定义为 []byte 数组
}


// 保存页面的结构体方法
func (p *Page) savePage() error { // 会返回一个错误值
	filename := p.Title + ".txt" // 以【标题】作为文件名
	return os.WriteFile(filename, p.Body, 0600) // 0600 为八进制, 表示只有当前用户才有读写权限 => 🌟 如果save成功，它返回nil; 如果失败, 它返回错误信息, 如果 err != nil, 则表示有 ❌ 错误
}


// 读取保存后的页面数据的方法 (加载一个页面数据并返回一个指向该页面的指针, 以便后续可以使用或修改这个页面的数据)
func loadPage(title string) (*Page, error) { // (返回 *Page 跟 err)  => *Page 表示返回值是是一个指针 => 指向 Page 结构体的指针, 如下 👇
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil { // 处理错误, 比如文件不存在
		return nil, err // 如果 err != nil, 则表示有 ❌ 错误, 返回 nil
	}
	return &Page{Title: title, Body: body}, nil // 👈 创建一个新的 Page 实例, 并返回它的指针 🌟
}



// 🌟 view 路由, 在此路由内通过 loadPage 读取页面内容, 并返回 html 页面
// 访问 http://localhost:8080/view/foo
func viewHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println("有人访问了这个地址: ", r.URL.Path) // 打印日志, 说明有人访问了这个地址

	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title) // _ 忽略错误

	fmt.Fprintf(w, "<h1>%s<h1> <p>%s</p>", p.Title, string(p.Body)) // 🔥 %s专门用于字符串 => ⚠️ 注意 p.Body是一个字节切片（[]byte）, 在将其作为字符串格式化时, 需要先将它转换成字符串
}



func main() {
	// 启动 http 服务, 访问 http://localhost:8080/view/foo
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}