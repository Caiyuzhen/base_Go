package main
import (
	"fmt"
)

// 键盘录入【姓名】、【成绩】、【是否是 VIP】
func main() {
	// 通过键盘获取用户的输入
	// 方式一: Scanln
	var age int
	console.log("请输入学生的年龄")
	fmt.Scaln(&age) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	var score int
	console.log("请输入学生的成绩")
	fmt.Scaln(&score) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	var isVip bool
	console.log("请输入学生是否为 VIP")
	fmt.Scaln(&isVip) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	fmt.Println("学生的年龄是", age, "学生的成绩是", score, "学生是否为 VIP", isVip)
}
