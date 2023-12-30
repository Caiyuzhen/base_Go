package main

import (
	"fmt"
	"strconv"
	"time"
)


func main() { 

	// 🔥【并发】多个协程
	for i := 1; i <= 5; i ++ {
		go func(x int) {  // 🌟 匿名协程函数
			fmt.Printf("🌟 匿名协程函数 - %v \n", strconv.Itoa(x)) // 因为是并发的, 🔥 所以不会按照顺序打印!!
		}(i) // 🔥 匿名函数的【传参】, i 形成闭包参数
	}


	time.Sleep(time.Second * 2) // 🌟 休眠两秒，防止主协程提前结束 ⚠️⚠️
	
}

