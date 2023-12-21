package main
import (
	"fmt"
)

// 键盘录入【姓名】、【成绩】、【是否是 VIP】
func main() {
	// 通过键盘获取用户的输入
	// 方式一: Scanln => 在换行时结束输入
	var age int
	fmt.Println("请输入学生的年龄")
	fmt.Scanln(&age) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	var score int
	fmt.Println("请输入学生的成绩")
	fmt.Scanln(&score) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	var isVip bool
	fmt.Println("请输入学生是否为 VIP")
	fmt.Scanln(&isVip) // 🚀 传入 &age 的地址这样才能改变 age 真实的值！！

	fmt.Println("学生的年龄是:", age, "学生的成绩是:", score, "学生是否为 VIP:", isVip)



	// fmt.Println("————————分割线————————")


	// 方式二: Scanf => 通过指定格式进行输入
	var name string
	var score_2 float64
	var isVip_2 bool
	fmt.Println("请输入学生的姓名、成绩、是否为 VIP, 用空格进行分割")
	fmt.Scanf("%s %f %t", &name,&score_2,&isVip_2) // %s %f %t 表示输入的是 字符串、浮点数、布尔值
	fmt.Println("学生的姓名是:", name, "学生的成绩是:", score_2, "学生是否为 VIP:", isVip_2)
}
