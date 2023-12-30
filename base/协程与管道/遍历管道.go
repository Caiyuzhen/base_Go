package main

import (
	"fmt"
)

/* 
	管道的遍历:
		1. 🔥 只能用【for-range】 来遍历管道, 不能用 for 循环, 因为管道的底层是【队列】而不是索引, 所以不能用 for 循环
		2. 🔥 在遍历时, 如果管道没有关闭, 则会报错 deadlock
		3. 🔥 在遍历时, 如果管道已经关闭, 则会正常遍历数据, 遍历完后, 会退出遍历
*/


func main() {
	var numChannel chan int = make(chan int, 100) // 🔥 声明管道

	// 往管道内放入 100 个元素 (放入数据可以用 for 循环)
	for i := 0; i < 100; i++ {
		numChannel <- i 
	}

	// 关闭管道 (🔥在遍历管道前没关闭的话, 遍历管道会报错 deadlock)
	close(numChannel)

	// 遍历管道 (因为管道底层是队列, 所以👇下面只能接到 value, 而接不到索引!!) => 🚀 遍历前记得关闭管道 !!!
	for value := range numChannel {
		fmt.Printf("管道的值: %v \n", value) // 打印出 0 - 99
	}

}

