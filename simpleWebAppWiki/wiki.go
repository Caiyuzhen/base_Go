package main

import (
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

// 【数据结构】定义 Pagr 的数据结构, 用于存储页面的标题和内容
type Page struct {
	Title string
	Body  []byte // 为了方便存储, 不定义为 string, 而是定义为 []byte 数组
}


// 【🌟结构体的方法 - 保存数据】保存页面
func (p *Page) savePage() error { // 会返回一个错误值
	filePath := "data/" + p.Title + ".txt"                // 以【标题】作为文件名
	return os.WriteFile(filePath, p.Body, 0600) // 0600 为八进制, 表示只有当前用户才有读写权限 => 🌟 如果save成功，它返回nil; 如果失败, 它返回错误信息, 如果 err != nil, 则表示有 ❌ 错误
}


// 【🌟结构体的方法 - 读取数据】读取保存后的页面数据 (加载一个页面数据并返回一个指向该页面的指针, 以便后续可以使用或修改这个页面的数据)
func loadPageData(title string) (*Page, error) { // (返回 *Page 跟 err)  => *Page 表示返回值是是一个指针 => 指向 Page 结构体的指针, 如下 👇
	filePath := "data/" + title + ".txt" // 比如读取 foo.txt 文件
	body, err := os.ReadFile(filePath)
	if err != nil { // 处理错误, 比如文件不存在
		return nil, err // 如果 err != nil, 则表示有 ❌ 错误, 返回 nil
	}
	return &Page{Title: title, Body: body}, nil // 👈 创建一个新的 Page 实例, 并返回它的指针 🌟
}





// 🌟 view 视图路由, 在此路由内通过 loadPageData 读取页面内容, 并返回 html 页面 __________________________________________________________________
// 访问 http://localhost:8080/view/foo => 页面内点击 edit 页
func viewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("有人访问了这个地址: ", r.URL.Path) // 打印日志, 说明有人访问了这个地址

	// 👇 加载数据, 放入结构体 => 把结构体内的数据 => 渲染到 html 上
	title := r.URL.Path[len("/view/"):]
	p, error := loadPageData(title) // _ 忽略错误 => 🌟 loadPageData 函数会返回 (*Page, error) , 分别是【页面内容】和【错误信息】viewHandler
	if error != nil {
		// http.NotFound(w, r)                                          // 如果页面不存在, 则返回 404 状态码
		http.Error(w, error.Error(), http.StatusInternalServerError) // 如果有错误, 则返回 500 状态码
		return
	}

	// 【方法一】使用 html 模板渲染页面
	readTemplate(w, "viewPage", p) // 使用抽象出来的读取方法 !
	// t, _ := template.ParseFiles("viewPage.html") // 解析 html 模板文件
	// t.Execute(w, p)

	// 【方法二】硬编码 html
	// fmt.Fprintf(w, "<h1>%s</h1> <p>%s</p>", p.Title, string(p.Body)) // 🔥 %s专门用于字符串 => ⚠️ 注意 p.Body是一个字节切片（[]byte）, 在将其作为字符串格式化时, 需要先将它转换成字符串
}





// 🌟 edit 视图路由 __________________________________________________________________________________________________
// 在 // 访问 http://localhost:8080/edit/foo 
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] // 获取页面标题（用标题读取页面.txt文件）) => len("/view/")将会返回6，因为字符串"/view/"有6个字符, 所以切片从第 6 个字符开始截取
	p, error := loadPageData(title)
	if error != nil {
		p = &Page{Title: title} // ⚠️ 如果页面不存在, 则创建一个新的 Page 实例
	}

	// 【方法一】使用 html 模板渲染页面
	readTemplate(w, "editPage", p) // 使用抽象出来的读取方法 !
	// t, _ := template.ParseFiles("editPage.html") // 解析 html 模板文件
	// t.Execute(w, p)  // w 表示一个 【* Page】 指针，指向的是要渲染到页面上的数据。在这个例子中，模板将使用 【p.Title】和 【p.Body】这样的字段来渲染HTM                           // 将 p 作为数据传递给模板, 并执行模板

	// 【方法二】硬编码 html
}




// 抽象出读取 html 模板的方法
func readTemplate(w http.ResponseWriter, tempName string, p *Page) {
	t, error := template.ParseFiles("html/" + tempName + ".html")
	// 打印模板加载路径
	log.Println("模板加载路径: ", "html/" + tempName + ".html")

	if error != nil {
		http.Error(w, "模板加载错误", http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}



// 🌟 save 编辑数据层的路由, 把 edit 页提交上来的路由进行储存 __________________________________________________________________________________________________
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]          // 从 URL 中获取页面标题
	body := r.FormValue("body")                  // 从 r 中获取 body 的值(对方会发送一个表单类型的 http 请哦去)
	p := &Page{Title: title, Body: []byte(body)} // 创建一个 Page 实例
	p.savePage()                                 // 保存页面 => 🔥 调用结构体内的方法
	fmt.Fprintf(w, "页面数据 %s 保存成功! ", title)   // 返回一个成功信息
	// 重定向回 view 页面
	// http.Redirect(w, r, "/view/" + title, http.StatusFound) // 重定向到 view 页面, StatusFound 表示 302 状态码
}



func main() {
	// 启动 http 服务, 访问 http://localhost:8080/view/foo
	http.HandleFunc("/view/", viewHandler) // 渲染 html 或 txt 页面 localhost:8080/view/foo
	http.HandleFunc("/edit/", editHandler) // 编辑页面
	http.HandleFunc("/save/", saveHandler) // 保存页面
	log.Fatal(http.ListenAndServe(":8080", nil))
}
