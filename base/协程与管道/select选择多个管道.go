package main

import (
	"fmt"
	"time" // 暂停使用 time 包
	_"sync"
)


/* 
	多个管道时候, 可以使用 select 来【随机公平】的选择管道
*/


func main() {
	// 管道 A - int 类型
	intChan := make(chan int, 10)
	go func(){ // 匿名函数
		time.Sleep(time.Second * 3) //暂停 3s
		intChan <- 10
	}()

	
	// 管道 B - string 类型
	stringChan := make(chan string, 10)
	go func() { // 匿名函数
		time.Sleep(time.Second * 1) //暂停 1s
		stringChan <- "Hello"
	}()

	// fmt.Println(<-intChan) // 取数据是阻塞的, 要等待所有管道都执行完后才会执行


	// 🌟 select 取数据
	select { // 随机公平, 因为 👆 上方的 intChain 的阻塞时间比 stringChan 的阻塞时间长, 🔥 所以会先执行 stringChan !!!
		case v := <- intChan:
			fmt.Println("intChan 取到数据", v)
		case v := <- stringChan:
			fmt.Println("stringChan 取到数据", v)
		default: // 🔥防止 select 自身被阻塞
			fmt.Println("默认数据, 防止 select 自身被阻塞")
	}
}
