package main
import "fmt"

// 在 Go 语言中没有 while 和 do while 循环，只有 for 循环。


// 主函数
func main() {
	// 循环求和功能

	// 🌟 一般 For 循环 ————————————————————————————————————————————————————————————————————————
	// 写法一
	sum := 0 // 🔥 for 循环的【初始化】不能用  var sum, int = 0 !!!

	for x := 1; x <= 6; x++ { // 🔥 for 循环的【初始化】不能用  var x int = 1 !!!
		sum += x
	}
	fmt.Println(sum)



	// 写法二(for):
	sum2 := 0
	zz := 0
	for zz <= 6 {
		sum2 += zz
		zz++
	}
	fmt.Println(sum2)


	// 🌟 For range 循环 => 这种方法可以取出【中文】 => 可以对数组、切片、字符串、map、通道 进行遍历 ————————————————————————————————————————————————————————————————————————
	// var content2 string = "hello golang"
	var content2 string = "hello 你好"
	for i , value := range content2 {
		fmt.Printf("索引为: %d, 具体的值为: %c \n", i, value) // ，每个中文占用 3 个字节 !!
	}


	// 一般 For 循环
	var content1 string = "hello C++" // ⚠️注意, 如果有中文的话, 需要对中文字符做一个切片
	for kk := 0; kk < len(content1); kk++ {
		fmt.Printf("%c", content1[kk]) // 普通 for 循环是按每个字符一个个进行输出的
	}

}
