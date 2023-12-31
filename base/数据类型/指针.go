package main
import (
	"fmt"
)

/* 
	指针:
		🚀 * 是用于声明指针类型, 也是用于指针取值 🚀
		& 是用于取地址

		🔥🔥 指针接收的一定是一个【内存地址】🔥🔥


*/


func main() {
	// 🔥 内存地址 ————————————————————————————————————————————————————————————————
	// 【基本变量】 & 来获取内存地址
	var age int = 18
	fmt.Println(&age)

	// 【指针变量】, ptr 表示指针变量的名称 | 类型是 *int => 指针类型(指向 int 类型的指针) | 值为 &age 这个地址
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr 这个指针变量的内存地址为:", &ptr)

	// 🔥 【通过指针改变值】 ————————————————————————————————————————————————————————————————
	var num int = 10
	var ptrr *int = &num //🔥🔥 指针接收的一定是一个【内存地址】🔥🔥!! 并可类型得对应!! 比如不能把 *int 的指针赋值给 &str 类型的地址!!
	*ptrr = 20
	fmt.Println("通过指针改变后的指:", num)


	a := 1
    XX(&a) // 传递 a 的地址
	fmt.Println(a) // 输出 2，a 被修改



	// 使用 new 函数直接定义一个指针变量
	numX := new(int) // 返回了一个指针 => *int 指针 (指向新分配的零值的指针)
	fmt.Printf("num 的类型是 %T, num 指针变量的值是 %v, num 指针的地址 %v, num 指针地址指向的值是: %v \n", numX, numX, &numX, *numX)

	// 更改 num 指针指向的内存地址的值
	*numX = 100
}




func XX(b *int) {
    *b = 2 // 修改指针指向的原始值
}