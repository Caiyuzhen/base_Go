package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
	🔥 Go 的并发都是通过【协程】来实现的

	🌟 协程的概念:
		协程是一种轻量级的【线程】, 也叫【用户态线程】或【绿色线程】

		协程是【函数级别】的操作, 因此创建和销毁的代价都比较小, 而以前的【线程】是【系统级别】的操作, 创建和销毁的代价比较大

		协程本质还是在【单个线程】内进行多个任务的切换（趁着 IO 阻塞的时候切换）
*/

func main() { // 👈 main 就是 主线程
	/* 
		Case: 主线程与协程并发执行
		1.协程打印 Hello
		2.主线程打印 你好
	*/
	go test() // 👈 协程函数, 加个 go 就是创建一个协程 🔥

	for i := 1; i <= 10; i++ {
		fmt.Println("Hello + ", strconv.Itoa(i)) // strconv.Itoa() 将 int 转换为 string

		// 阻塞一秒
		time.Sleep(time.Second) // 每个一秒主线程都会阻塞一次
	}

	
}


// 创建【协程函数】 => 主死从随（🔥主线程停止了, 协程也会停止）
func test() {
	for i := 1; i <= 1000; i++ {
		fmt.Println("你好 + ", strconv.Itoa(i)) // strconv.Itoa() 将 int 转换为 string

		// 阻塞一秒
		time.Sleep(time.Second) // 每个一秒协程都会阻塞一次
	}
}