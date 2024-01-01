package main
import ("fmt")


/* 
	What?
		defer 语句会将函数押入到一个栈中，当函数执行完毕后，再从栈中取出执行, 可以利用此机制来进行错误处理
	
	How?
		defer + 匿名函数捕捉错误
*/

func main() {
	test()
	fmt.Println("正常运行到此处, 不会被错误所打断")
}


func test() {
	// 🌟 调用匿名函数来专门捕获错误
	defer func() { 
		err := recover() // // 调用 recover() 函数可以捕获到 panic 异常
		if err != nil { // // 如果不等于空, 则表示捕捉到错误
			fmt.Println("已经捕获错误, 错误为: ", err)
		}
	}()
	num1 := 10
	num2 := 0 
	result := num1 / num2 // 除数不能为 0 , 逻辑上会报错
	fmt.Println(result)
}