package main
import "fmt"


// 开始执行的入口函数
func main() {
	name := "Jimmy" // 🚀省略 var 并且自动类型推断
	var name2 string = "jimmy2" // 手动定义类型
	age, address, hasHobby := 12, "NY", true // 定义多个数据并进行自动推断

	// 数组
	myArr := [3]string{} // 使用自动类型推断和初始化数组 {}
	myArr[0] = "😄"
	myArr[1] = "💔"
	myArr[2] = "😭"

	// 映射
	myMap := make(map[string]string)  // 使用 make 初始化映射
	myMap["robot"] = "🤖️"
	myMap["player"] = "🎮"
	myMap["rocket"] = "🚀"

	// 遍历
	for x := 0; x < 3; x++ {
		fmt.Println("⚡️")
	}

	// 条件判断
	animal := "🐱"
	if animal == "🐱" {
		fmt.Println("😍")
	} else {
		fmt.Println("😭")
	}

	// 指针
	var year int = 2024
	var p *int = &year // & 是取地址符号, 用来获取 year 变量的地址 | p 是指针变量, * 是一个解引用符号, 用来获取指针指向的变量的值, 🔥 因此这里把 p 定义为了 【*int】 整型变量, 然后指向了存储了 year 变量的内存地址
	// 可以通过 *p 来修改 year 的值
	fmt.Println(*p)
	*p = 2025

	// 打印单行
	fmt.Println("Hi~", name, name2, age, address, hasHobby, p, *p)

	// 打印多行
	fmt.Println(`aaaaabbbbbcccccdddddeeeefff
		aaaaabbbbbcccccdddddeeeefffaaaaabbbbbcccc
		cdddddeeeefff`)
}