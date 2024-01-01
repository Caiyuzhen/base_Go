package main
import (
	"fmt"
	"net" // 网络通信包
)

/* 
	Net 包:
		客户端通过 Dial 函数来建立连接

		服务端通过 Listen 函数来监听端口
*/


// 4. 定义处理【客户端】发送过来的数据的协程
func process(connect net.Conn) { // 👉 net.Conn 为 connect 的类型~~~
	// 🚀 连接用完一定要关!!
	defer connect.Close()

	for {
		// 🔥🔥创建切片, 把读取到的数据放入切片中
		buf := make([]byte, 1024) // 🌟 make => 创建切片

		// 从 connect 中读取数据, 放入 buf 中, 返回 读取到的字节数 | 错误
		dataNum, err := connect.Read(buf)
		if err != nil {
			fmt.Println("❌ 读取失败...", err)
			return
		}

		// 将读取的内容输出在【服务器】的控制台内
		res := string(buf[0:dataNum]) // 🔥 把 buf 这个切片转换为字符串, 🔥 范围为 0 到 dataNum 字节
		fmt.Printf("👀 读取到数据为: %v", res)
	}
}



func main() {
	// 1. 创建服务器
	fmt.Println("🚀 服务器启动...")


	// 2.创建监听套接字
	listenen, err := net.Listen("tcp", "192.168.0.118:3030") // 🌟 调用 Listen => 参数为 tcp 协议 | 端口
	if err != nil {
		fmt.Println("❌ 监听失败...", err)
		return
	}
	fmt.Println("✅ 监听成功...", listenen)


	// 3.🔥轮询等待客户端的连接
	for {
		connect, err := listenen.Accept() // 🌟 调用 Accept => 等待客户端连接
		if err != nil {
			fmt.Println("❌ 客户端等待连接失败...", err)
			return
		} else {
			fmt.Printf("✅ 客户端等待连接成功! con: %v, \n 客户端的 id 地址为: %v \n", connect, connect.RemoteAddr().String()) // 🌟 connect.RemoteAddr().String() 可以获取客户端的 ip 地址
		}

		// .5 🔥开启协程来处理客户端的请求
		go process(connect) // 🔥 不同的【客户端连接】, 会开启不同的协程处理
	}
}
