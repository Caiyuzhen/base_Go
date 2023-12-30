package main

import (
	"fmt"
)

/* 
	管道 channel :
		1. 本质上就是一个队列型的数据结构
		2. 管道是先进先出, 如果是栈的话, 后进先出
		3. 管道自身是【线程安全】的, 多协程访问时, 不需要加锁
		4. 管道是有类型的, 一个整数管道只能存放整数, 一个字符串管道只能存放字符串
		5. 管道是引用类型, 必须初始化才能使用 🔥
*/	



func main() {
	// 🔥 声明管道
	var numChannel chan int = make(chan int, 10) // 10 代表管道的容量, 10 个元素

	
	// 向管道内写入数据 => 管道（channel）的数据是按照先进先出（FIFO）的顺序进行存取的。这意味着你不能直接访问管道中的某个特定元素，只能按照元素进入管道的顺序依次取出
	numChannel <- 10
	numChannel <- 20
	numChannel <- 30
	num := 99
	numChannel <- num // 将变量 num 写入管道


	// 从管道内读取数据, 读取是依次读取的 (🔥 在没有使用协程的情况下, 读取管道时, 必须保证管道中已经有数据, 否则会报错)
	res := <- numChannel


	fmt.Printf("管道的值： %v \n", res) // 10
	fmt.Printf("管道的下一个值： %v \n", <- numChannel) // 打印下一个值 20


	// 关闭管道 (🔥 关闭后不可写入, 但是可以读取)
	close(numChannel)

	res2 := <- numChannel
	fmt.Printf("读取管道内的数据: %v \n", res2) // 30

}