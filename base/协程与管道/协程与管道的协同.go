package main

import (
	"fmt"
	"time"
	"sync"
)

/*
	Case:
		一个写协程, 向管道内写入 50 个整数
		一个读协程, 从管道内写入的数据
		主线程等待读写协程都完成后才退出

	使用 协程 + 管道 + waitGroup 实现, 不需要用【锁】, 因为管道本身就是线程安全的:
		1. 同步机制: 
			管道内部实现了同步机制，确保了同时只有一个协程可以对管道进行发送或接收操作
		2. 阻塞和同步: 
			管道的发送和接收操作在必要时会阻塞协程
*/

var waitGroup sync.WaitGroup // 🔥🔥 防止主线程提前退出 !!

//写协程
func writeData(channel_data chan int) {
	for i := 0; i <= 50; i++ {
		channel_data <- i // 把 i 写入管道
		fmt.Println("写入数据: ", i)
		time.Sleep(time.Second) // 慢一点以便于观看, 写一次停顿一秒
	}
	close(channel_data) // 🚀 关闭管道
	defer waitGroup.Done() // 计数器 -1
}


// 读协程
func readData(channel_data chan int) {
	for value := range channel_data {
		fmt.Println("读取数据: ", value)
		time.Sleep(time.Second) // 慢一点以便于观看, 写一次停顿一秒
	}
	defer waitGroup.Done() // 计数器 -1
}


func main() {
	CHANNEL_DATA := make(chan int, 50) // 🔥 声明管道

	waitGroup.Add(2)

	go writeData(CHANNEL_DATA)
	go readData(CHANNEL_DATA)

	waitGroup.Wait() // 🚀🚀 阻塞主进程, 直到 waitGroup 为 0 为止
}