package main
import (
	"fmt"
	"crossPkgStruct/B"
)


//跨包创建结构体, 需要在跟路径上 go mod init crossPkgStruct
func main() {
	// 🔥 跨包创建结构体的实例 (方式一)
	var s model.Student = model.Student{"Kim", 29}
	fmt.Println(s)


	// 🔥 跨包创建结构体的实例 (方式二 - 类型推断)
	s2 := model.Student{"Annie", 22}
	fmt.Println(s2)


	// 🚀 使用工厂函数实例化结构体
	s3 := model.NewStudent("Zeno", 27)
	fmt.Println(*s3) // 👈 因为是指针, 所以用解引用来取值!!
}