package main

import (
	"bufio"
	"fmt"
	"net" // 网络通信包
	"os"
)

/*
	Net 包:
		客户端通过 Dial 函数来建立连接

		服务端通过 Listen 函数来监听端口
*/

func main() {
	// 1. 准备
	fmt.Println("😄 客户端启动...")


	// 2. 连接服务器
	connect, err := net.Dial("tcp", "192.168.0.118:3030") // 🌟 调用 Dial => 参数为 tcp 协议 | ip + 端口
	if err != nil {
		fmt.Println("❌ 连接失败...", err)
		return
	}
	fmt.Println("✅ 连接成功...", connect)


	// 3. 🔥 获取终端的输入并保存在【缓冲区】中 
	reader := bufio.NewReader(os.Stdin) // 🌟 bufio.NewReader => 从终端获取输入, 形成一个 reader 流对象
	for { // 🌟 循环读取终端输入
		// 4. 读取缓冲区内的内容
		fmt.Println("✏️ 请输入内容")

		str, err :=	reader.ReadString('\n') // 🌟 reader.ReadString => 读取字符串, 直到遇到 \n 结束
		if err != nil {
			fmt.Println("❌ 读取失败...", err)
			return
		}

		// ⚠️ 如果用户输入特定的退出命令，比如"exit"，则退出循环
		if str == "exit\n" {
			break
		}

		// 5. 🔥 发送数据给服务器, 参数需要传入【字节切片】, 因为 str 是字符串, 所以需要转换
		dataNum, err := connect.Write([]byte(str)) // 🌟 connect.Write => 发送数据给服务器 , 返回 发送的字节数 ｜ 错误
		if err != nil {
			fmt.Println("❌ 发送失败...", err)
			return
		}
		fmt.Printf("✅ 数据发送成功, 一共发送了 %v 字节的数据", dataNum) // data 表示发送的字节数 => 多少 KB
		fmt.Println("✏️ 请继续输入内容")
	}
}