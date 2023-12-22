package main
import ("fmt")

/* 
	🌟 在程序内只需要使用一次的函数, 可以使用匿名函数来定义
		比如定义时就调用
 */


func main() {
	// 🌟 Case1: 常用 => 定义匿名函数并进行调用
	res := func (num1 int, num2 int) int {
		return num1 + num2
	}(10, 20) 

	fmt.Println("res = ", res)


	// 🌟 Case2: 一般常用的 => 给变量赋值匿名函数调用
	fn := func (num1 int, num2 int) int {
		return num1 - num2
	}

	res2 := fn(10, 20)
	fmt.Println("res2 = ", res2)
}