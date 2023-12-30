package main

import (
	"fmt"
	"strconv"
	"sync"
)

/* 
	WaitGroup => 让【主线程】等待所有的【子线程】执行完毕再结束, 否则主线程结束后子线程也会结束了
*/

var waitGroup sync.WaitGroup  // 类似于计数器 => sync.WaitGroup是一个计数器, 用于等待指定数量的事件完成

func main() {

	// 【🚀写法 2 】🔥 启动 5 个计数器 🔥 => 总数是 5 个读 => 记得改 !!
	// waitGroup.Add(2)

	// 启动 5 个协程
	for i := 0; i < 5; i++ {
		// 🔥 【1 - 启动 x 个计数器】 加一, 表示新 + 1 个协程计数器
		waitGroup.Add(1) // 【🚀写法 1 】每次增加一个协程计数器 => 记得改!! 

		go func(i int) {
			fmt.Println("子线程" + strconv.Itoa(i) + "执行完毕")
			// 🔥 【2 - 计数器】 减一, 表示一个协程执行完毕 => 协程执行完成 -1
			defer waitGroup.Done() 
		}(i)
	}

	// 🚀 【3 - 阻塞】主线程 => 等待所有的子线程执行完毕 => 计数器为 0
	waitGroup.Wait()
}