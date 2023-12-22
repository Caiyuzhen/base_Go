package main
import ("fmt")

// 第一步: 🌟 最先全局变量的定义
var num int = test()

func test() int {
	fmt.Println("执行了 test 函数")
	return 250
}

// 第二步: 🌟 init 函数, 会在 main 函数执行之前就执行
func init() {
	fmt.Println("执行了 init 函数")
}

// 第三步: 执行 main 函数
func main() {
	fmt.Println("执行了 main 函数")
}