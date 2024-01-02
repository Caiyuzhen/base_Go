package main

import (
	// "fmt"
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