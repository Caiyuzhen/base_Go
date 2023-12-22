package main
import ("fmt")


// 求和函数
func getSum() func (int) int { // 返回值为  func (int) 且函数的参数是 int, 整体返回值是 int
	var sum int = 0
	return func(x int) int {  // 返回函数, 此时 sum 就被闭包了 => sum 就是一个局部变量了, 没有被释放内存
		sum = sum + x // => 闭包中使用的变量 🔥 会【!一直保存在内存!】中可以被使用!! 🔥 => 但是对内存消耗很大!!
		return sum
	}
}


func main() {
	f := getSum() // 🔥获得闭包函数!!
	res := f(1)
	fmt.Println("res = ", res) // 1

	res2 := f(1)
	fmt.Println("res2 = ", res2) // 2
}