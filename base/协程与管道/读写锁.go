package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	🔒 读写锁
		读写锁的效率比互斥锁高：
			互斥锁是完全互斥的,
			读写锁是读共享, 写互斥 => 都在读则不互斥, 一个读一个写才互斥

*/

var totalNumber int
var waitGroup sync.WaitGroup
var lock sync.RWMutex // 🔒 读写锁

// 协程函数 1
func read() {

	// 🔒 读取数据时, 加读锁 => 只读不影响, 一读一写才互斥上锁
	lock.RLock()

	fmt.Println("开始读取数据")
	time.Sleep(time.Second * 1)
	fmt.Println("读取数据完毕 ~ totalNumber = ", totalNumber)

	lock.RUnlock() // 解锁

	// 计数器 -1
	defer waitGroup.Done()
}

// 协程函数 2
func write() {

	// 🔒 写入数据时, 加写锁 => 读写都互斥上锁
	lock.Lock() // 默认就是写锁
	fmt.Println("开始写入数据")
	time.Sleep(time.Second * 5) // 模拟写入数据需要 10s
	totalNumber = 999
	fmt.Println("写入数据完毕 ~ totalNumber = ", totalNumber)

	lock.Unlock() // 解锁

	// 计数器 -1
	defer waitGroup.Done()
}

// 主线程
func main() {
	// 🔥 启动 6 个计数器 🔥 => 总数是 5 个读 + 1 个写!!! => 记得改 !!
	waitGroup.Add(6)

	// 启动协程 => 模拟读多写少的场景
	for i := 0; i < 5; i++ {
		go read() // 读 5 次
	}

	go write()

	// 阻塞主进程
	waitGroup.Wait()
}

