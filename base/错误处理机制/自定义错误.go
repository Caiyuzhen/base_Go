package main
import (
	"fmt"
	"errors"
)

/* 
	自定义错误
*/


func main() {
	err := test()
	if err != nil {
		fmt.Println("已经捕获错误, 错误为: ", err)

		// 🔥 抛出错误后且停止运行程序
		panic(err)  // 打印出 => panic: ❌ 除数不能为 0
	}

	fmt.Println("正常运行到此处, 不会被错误所打断")	
}



func test() (err error) { // 🔥 定义返回值为 err, 位 error 类型
	num1 := 10
	num2 := 0 

	// 🚀🚀 自定义捕获错误的机制
	if num2 == 0 {
		// 🌟 【】 抛出自定义错误
		return errors.New("❌ 除数不能为 0")
	} else {
		result := num1 / num2 // 除数不能为 0 , 逻辑上会报错
		fmt.Println(result)
		fmt.Println("✅ 除数不为 0, 正确")
		return nil // 没有错误则返回 0 值 => nil
	}
}