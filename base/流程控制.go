package main
import (
	"fmt"
	"math"
)

/*
	Go 条件分支的特点
		if、switch 后面可以放表达式（变量、常量、函数的等）
		switch 内的 case 类型必须和 switch 表达式的类型一致
		case 内可以放多个值, 比如 Case 1, 2, 3, ...
		可以不写 break
		default 可以不写, 也可以放在任意位置
		case 内可以类似 if-else 分支来使用, 比如 case age > 35:
		case 后不能用 var 来声明变量, 但是可以用 := 来声明变量, 但是需要用; 来结尾
		可以使用 fallthrough 来穿透一层
		
*/

func main() {
	// if 分支
	var age int
	age = 19
	if age > 18 {
		fmt.Println("你已经成年了")
	} else {
		fmt.Println("你还未成年")
	}


	// if 后方可以跟变量的定义, 【格式化】esle 不能换行写, 否则会报错!
	if count := 40; count >= 80 {
		fmt.Println("库存充足")
	} else if count > 30 && count <= 50 {
		fmt.Println("库存一般")
	} else {
		fmt.Println("库存不足")
	}


	// switch 分支
	var score float64 = 87

	// switch 后面可以跟表达式
	switch int(math.Floor(score/10)) { // math.Floor 向下取整, 否则 8.7 浮点数会匹配到不及格!
		case 10:
			fmt.Println("满分")
		case 9:
			fmt.Println("优秀")
		case 8:
			fmt.Println("良好")
			fallthrough // 向下穿透一层!
		case 6, 7:
			fmt.Println("中等")
		default:
			fmt.Println("不及格")
	}

	fmt.Println(score/10) // 8.7
	
}