package main
import (
	"fmt"
)




func main() {
	// 🔥 内存地址 ————————————————————————————————————————————————————————————————
	// 【基本变量】 & 来获取内存地址
	var age int = 18
	fmt.Println(&age)

	// 【指针变量】, ptr 表示指针变量的名称 | 类型是 *int => 指针类型(指向 int 类型的指针) | 值为 &age
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 这个指针变量的内存地址为:", &ptr)

	// 🔥 【通过指针改变值】 ————————————————————————————————————————————————————————————————
	var num int = 10
	var ptrr *int = &num //🔥🔥指针接收的一定是一个【内存地址】!! 并可类型得对应!! 比如不能把 *int 的指针赋值给 &str 类型的地址!!
	*ptrr = 20
	fmt.Println("通过指针改变后的指:", num)


}