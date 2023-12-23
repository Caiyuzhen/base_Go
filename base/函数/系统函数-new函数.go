package main
import (
	"fmt"
)


/* 
	🌟 new 函数 => 用来分配内存, 返回值是对应类型的【指针】 => 🔥可以让数据变成引用类型🔥
		可以用来分配 int 、 float 、 bool 、 string 、 struct 等基本类型的内存
		而其 指针 、 slice 、 map 、 切片 、 管道chan 、 interface 等都是引用类型, 则需要用 make 函数来分配内存


	🔥 跟 JS 的区别, JS 无法自定义传递【引用类型的数据】, 而 Go 可以自己定义传递【引用类型的数据】
		在 JS 中，原始数据类型（如数字、字符串、布尔值）总是通过值传递
		而对象（包括数组和函数）则是通过引用传递
*/

func main() {
	num := new(int) // 返回了一个指针 => *int 指针 (指向新分配的零值的指针)
	fmt.Printf("num 的类型是 %T, num 指针变量的值是 %v, num 指针的地址 %v, num 指针地址指向的值是: %v \n", num, num, &num, *num)

	// 更改 num 指针指向的内存地址的值
	*num = 100

	// 再次打印，观察改变后的值
	fmt.Printf("修改后, num 指针地址指向的值：%v\n", *num)
	/*
		🌟 【指针变量】跟【普通变量】各有各的内存地址, 但是【指针变量】的值是【普通变量】的内存地址
			num 的类型是 *int, 
			num 指针变量的值是 0x1400000e108, 
			num 指针的地址 0x14000050020, 
			num 指针地址指向的值是: 0 

			%T 格式化指令用于打印变量的类型。
			%v 格式化指令用于打印变量的值。
			& 用于获取变量的地址。
			* 用于从指针访问目标对象（即指针解引用），这里用于更改和访问num指针指向的内存地址的值。
	*/
}