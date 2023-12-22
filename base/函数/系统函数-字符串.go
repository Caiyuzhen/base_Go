package main
import (
	"fmt"
	"strconv"
	"strings"
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
	fmt.Println("————————————————————————————————————")


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


	// 统计字符串内有多少个指定的子串 => 比如 Go 在 GolangGoGoGo 内出现了几次
	res4 := strings.Count("GolangGoGoGo", "Go")
	fmt.Println(res4) // 4


	// 不区分大小写的字符串比较 => 🌟 可以用于校验密码
	res5 := strings.EqualFold("abc", "Abc")
	fmt.Println(res5) // true


	// 区分大小写的字符串比较
	strA := "abc"
	strB := "Abc"


	// 使用 == 操作符进行比较 => 🌟 可以用于校验密码
	isEqual := strA == strB
	fmt.Println("区分大小写的比较结果: ", isEqual) // false


	fmt.Println("————————————————————————————————————")


	// 返回子串在字符串第一次出现的 index 值, 如果没有返回 -1
	index := strings.Index("NLT_abc", "abc")
	fmt.Println(index) // 4


	// 字符串的替换 => 比如把 go 替换为 ab
	strX := "golangAndJAVAvsPythonAndgoAndSuperGoAndFastGo"
	strY := strings.Replace(strX, "go", "C++", -1) // -1 表示全部替换
	strZ := strings.Replace(strX, "go", "C++", 2) // 2 表示替换前两个
	fmt.Println(strY) // C++langAndJAVAvsPythonAndC++AndSuperC++AndFastC++
	fmt.Println(strZ) // C++langAndJAVAvsPythonAndC++AndSuperGoAndFastGo


	// 按照指定的某个字符, 为分割标识, 将一个字符串拆分成字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	fmt.Println(strArr) // [hello world ok]


	// 将字符串的字母进行大小写的转换
	strU := "golangAndJAVAvsPythonAndgoAndSuperGoAndFastGo"
	strL := strings.ToLower(strU)
	strUp := strings.ToUpper(strU)
	fmt.Println(strL) // golangandjavavspythonandgoandsupergoandfastgo
	fmt.Println(strUp) // GOLANGANDJAVAVSPYTHONANDGOANDSUPERGOANDFASTGO


	// 将字符串左右两边的空格去掉
	strTrim := strings.TrimSpace("  hello world  ")
	fmt.Println(strTrim) // hello world


	// 将字符串左右两边指定的字符去掉
	strTrim2 := strings.Trim("! hello world !", " !")
	fmt.Println(strTrim2) // hello world


	// 将字符串左边指定的字符去掉
	strTrim3 := strings.TrimLeft("! hello world !", " !")
	fmt.Println(strTrim3) // hello world !


	// 将字符串右边指定的字符去掉
	strTrim4 := strings.TrimRight("! hello world !", " !")
	fmt.Println(strTrim4) // ! hello world


	// 判断字符串是否以指定的字符串开头 => 🌟 比如判断是不是以 http 开头
	isPrefix := strings.HasPrefix("http://java.fun.com/jsp/shtl/fmt", "http")
	fmt.Println(isPrefix) // true


	// 判断字符串是否以指定的字符串结尾 => 🌟 比如判断是不是以 .jpg 结尾
	isSuffix := strings.HasSuffix("demo.jpg", ".png")
	fmt.Println(isSuffix) // false
}