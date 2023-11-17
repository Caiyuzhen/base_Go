package main
import "fmt"

// 一次性声明多个全局变量
var (
	n8 = 999
	n9 = 666
)


func main() {
	// 第 1 种.声明 + 类型 + 赋值
	const Pi float64 = 3.14 // Go 的浮点数一般使用 float64, 会更精确
	var age int = 88


	// 第 2 种.声明 + 类型
	var name string
	name = "Kim"


	// 第 3 种.声明 +  赋值(自动推断类型)
	var isTop = true

	// 第 4 种.省略 var 或 const 声明, 用 := 来自进行动推断类型
	isBoy := false

	// 声明多个变量
	var n1, n2, n3 int // 默认值为 0

	height, weight := 177.8, 62.4 // 声明 + 赋值 + 自动类型推断

	fmt.Println(Pi, age, name, isTop, isBoy, n1, n2, n3, height, weight, n8, n9)
}