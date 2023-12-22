package main
import ("fmt")

// 🌟 在数据库连接后需要释放资源, 此时可以用 defer 关键字 => 用于延迟一个函数或者方法的执行

func main() {
	fmt.Println(add(30, 60)) // 【第 4 执行】
}

func add(num1 int, num2 int) int {
	// 在 Go 中程序遇到 defer, 则不会立即执行 defer 后的语句, 而是把 defer 押入栈中, 然后继续执行后面的语句
	// 栈的特点是【先进后出!!】
	defer fmt.Println("num1=", num1) // => 放入 defer 栈, 后执行 【第 3 执行】 30 => 🌟 会保存押入栈时的值!! 可以延时使用这个值!!
	defer fmt.Println("num2=", num2) // => 放入 defer 栈, 先执行 【第 2 执行】 60 => 🌟 会保存押入栈时的值!! 可以延时使用这个值!!

	num1 += 90 // 120
	num2 += 50 // 110

	var sum int = num1 + num2
	fmt.Println("sum=", sum) // 【第 1 执行】 230
	return sum
}