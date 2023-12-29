package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	// 打开文件

	file, err := os.OpenFile("./file/Demo.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)  // 传入地址, 读写 or 追加数据 or 创建文件, 权限

	if err != nil { // 文件打开失败
		fmt.Println("open file err = ", err)
		return
	}

	// 关闭文件
	defer file.Close()

	// 把程序中的数据写入到文件内
	// IO 流 => 缓冲输出流(带缓冲区, 先写入缓冲区, 再写入文件效率比较高) => 文件
	writer := bufio.NewWriter(file) // 创建一个带缓冲的写入器
	writer.WriteString("Hello, 这是要写入的数据\n") // 把数据先写入到【缓存区】
	writer.Flush() // 把【缓冲区】的数据【真正写入到文件】

}