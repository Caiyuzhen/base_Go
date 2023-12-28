package main
import (
	"fmt"
	// "os"
	"os"
)

/* 
	IO 流
		流:
			text.txt => 程序 => 目标文件

		- 输入流: 从外部读取数据到程序中  
		- 输出流: 从程序输出数据到外部
*/

func main() {
	// 读取文件
	content, err := os.ReadFile("./file/text.txt") // 返回【文件内容 []byte】和【错误信息】
	if err != nil {
		fmt.Println("read file err = ", err)
	} else {
		fmt.Printf("content = %v", string(content)) // 🔥 打印的是【文件内容】
	}
}