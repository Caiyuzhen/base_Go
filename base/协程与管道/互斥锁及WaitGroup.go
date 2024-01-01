package main

import (
	"fmt"
	"sync"
)

/* 
	⏰ WaitGroup:
		防止主进程提前退出, 从而导致协程还没执行完毕就退出了
		🌟 在主进程中, 调用 wg.Add(2) 来增加计数器, 2 表示 2 个协程, 可以写 N 个

	🔒 互斥锁:
		当多个协程操作同个数据, 会出现数据不一致的情况, 这时候就需要用到互斥锁
*/


var totalNumber int
var waitGroup sync.WaitGroup
var lock sync.Mutex // 🔒 互斥锁

// 协程函数 1
func add() {

	for i := 0; i < 1000; i++ {
		// 🔒 修改数据时, 加锁
		lock.Lock()
		totalNumber = totalNumber + 1
		// 🔒 修改数据完毕, 解锁
		lock.Unlock()
	}

	// 计数器 -1
	defer waitGroup.Done()
}

// 协程函数 2
func reduce() {

	for i := 0; i < 1000; i++ {
		// 🔒 修改数据时, 加锁
		lock.Lock()
		totalNumber = totalNumber - 1
		// 🔒 修改数据完毕, 解锁
		lock.Unlock()
	}

	// 计数器 -1
	defer waitGroup.Done()
}

// 主线程
func main() {
	// 🔥 启动 2 个计数器 🔥 => 总数是 1 个读 + 1 个写!!!  => 记得改 !!
	waitGroup.Add(2)

	// 启动协程
	go add()
	go reduce()

	// 🚀🚀 阻塞主进程, 直到 waitGroup 为 0 为止
	waitGroup.Wait()

	// 打印
	fmt.Println("totalNumber = ", totalNumber) // 加了锁的逻辑, 结果一定是 0
}
