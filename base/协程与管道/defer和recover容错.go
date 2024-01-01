package main
import (
	"fmt"
	"sync"
)

/* 
	多个协程同时工作时, 如果有一个协程崩了, 就会发生 panic
	为了防止 panic, 可以使用 defer + recover 来捕获 panic
*/

var wg sync.WaitGroup

// 输出数字
func outPut() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	defer wg.Done()
}


// 除法操作
func devide() {
	defer wg.Done() // 计数器 -1 得放到【🔥第一位】, 否则下面捕捉错误后, 无法执行到这里, 就会发生死锁!!

	// defer + 匿名函数捕捉错误
	defer func() {
		err := recover() // 捕捉错误
		if err != nil { // 如果不等于空, 则表示捕捉到错误
			fmt.Println("已经捕获错误, 错误为: ", err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2 // 模拟报错 => 除数不能为 0
	fmt.Println(result)
}


func main() {
	wg.Add(2) // `⏰ WaitGroup: 防止主进程提前退出, 从而导致协程还没执行完毕就退出了`
	go outPut()
	go devide()
	wg.Wait() // 🚀🚀 阻塞主进程, 直到 wg 为 0 为止
}