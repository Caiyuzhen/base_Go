package main

import (
	"fmt"
	"time" // 暂停使用 time 包
	// _"time" // 暂停使用 time 包
	"sync"
)

/*
	死锁:
		当管道只写入而不读取时, 就会产生死锁

	如读写频率不一致的情况:
		如果是【写的快】但【读得慢】, 此时不会【死锁】
		如果是 【写的慢】但 【读得快】, 此时也不会出现【死锁】
*/

var waitGroup sync.WaitGroup // 🔥🔥 防止主线程提前退出 !!

//写协程
func writeData(channel_data chan int) {
	for i := 0; i <= 10; i++ {
		channel_data <- i // 把 i 写入管道
		fmt.Println("写入数据: ", i)
		time.Sleep(time.Second) // 模拟写的慢, 读得快
	}
	close(channel_data) // 🚀 关闭管道
	defer waitGroup.Done() // 计数器 -1
}


// 读协程
func readData(channel_data chan int) {
	for value := range channel_data {
		fmt.Println("读取数据: ", value)
		// time.Sleep(time.Second) // 模拟写的快, 读得慢
	}
	defer waitGroup.Done() // 计数器 -1
}


func main() {
	CHANNEL_DATA := make(chan int, 10) // 🔥 声明管道

	waitGroup.Add(2)

	go writeData(CHANNEL_DATA)
	go readData(CHANNEL_DATA)

	waitGroup.Wait() // 🚀🚀 阻塞主进程, 直到 waitGroup 为 0 为止
}