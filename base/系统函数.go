package main
import (
	"fmt"
	"strconv"
)


/* 
	字符串函数
		1. len(str) => 求长度
		2. 字符串遍历, 同时处理有中文的问题 r := []rune(str)
*/
func main() {
	// len(str) => 求长度
	str := "hello"
	fmt.Println(len(str)) // 5


	// 对字符串进行遍历（方法一) - for ... range
	for i, value := range str {
		fmt.Printf("索引为: %d, 值为: %c \n ", i, value)
	}

	fmt.Println("————————————————————————————————————·")

	// 对字符串进行遍历（方法二) - 切片方式
	r := []rune(str) // Go 数组的长度不可改变, 在特定场景中这样的集合就不太适用 => Go中的切片是一个引用类型, 它可以更加的灵活, 自定义长度
	for x := 0; x < len(r); x++ {
		fmt.Printf("索引为: %d, 值为: %c \n ", x, r[x])
	}

	// 类型转换（字符串 => 转整数）
	var str2 string = "123"
	res, err := strconv.Atoi(str2) // 🌟 => 转为 123, 同时如果有 err 可以接收 err
	res2, _ := strconv.Atoi(str2) // 🌟 => 转为 123, 可以忽略 err 
	fmt.Println(res, err, res2)

	// 类型转换（整数 => 转字符串）
	var num1 int = 123
	res3 := strconv.Itoa(num1)
	fmt.Println(res3)
}