package main

import (
	"fmt"
)


/* 
	只读只写只是管道的【性质】而非【类型】!

	1. 只读通道通常用于限制函数的能力，确保该函数只能从通道读取数据，而不能向其发送数据
	2. 只写通道通常用于限制函数的能力，确保该函数只能向通道发送数据，而不能从中读取数据
*/


func main() {
	// 👀 ✏️ 默认声明的管道为双向管道（读写）
	channel_data := make(chan int, 50) // 🔥 声明管道
	fmt.Println("channel_data: ", channel_data)



	// ✏️ 声明单向管道（只写, 加个箭头）
	// 只写声明方式一:
	var channel_data2 chan <- int
	channel_data2 = make(chan int, 3)
	channel_data2 <- 190
	fmt.Println("channel_data2: ", channel_data2)

	// ✏️ 只写声明方式二:
	channel_data3 := make(chan <- int, 3) // 创建一个缓冲大小为3的只写通道
	channel_data3 <- 190
	fmt.Println("channel_data3: ", channel_data3)
	// num := <- channel_data3 // 报错, 因为不可读，只能写



	// ✏️ 声明单向管道（只读, 加个箭头）
	// 只读声明方式一:
	var channel_data4 <- chan int
	channel_data4 = make(chan int, 3)
	fmt.Println("channel_data4: ", channel_data4)

	// 只读声明方式二:
	channel_data5 := make(<- chan int, 3) // 创建一个缓冲大小为3的只读通道
	fmt.Println("channel_data5: ", channel_data5)
	// channel_data5 <- 190 // 报错, 因为不可写，只能读




}