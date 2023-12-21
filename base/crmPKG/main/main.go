package main // 建议和所在的文件夹名字保持一致
import (
	"fmt"
	"crmPKG/dbUtils"
	otherName "crmPKG/data" // 使用别名定义包!!
)


// 👇 main 函数一定要放在引入的 package main 包下 !!!
func main() {
	fmt.Println("访问到了: ", otherName.StuNo) // 用别名来引用包内的变量!!
	fmt.Println("这是 main 函数~")
	dbUtils.GetConnent() // 🌟 调用 dbUtils 包下的 GetConnent 函数, 需要写文件夹名称 !!
	dbUtils.Calculate()
}