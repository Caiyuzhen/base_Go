package main
import ("fmt")


/* 
	空接口可以认为所有类型都继承了此接口, 因此可以接收任意类型的参数
*/



type T interface {

}

type Student struct {
	Name string
}




func main() {
	studentA := Student{"Kim"}

	var e T = studentA // 【写法一】空接口可以接收任意类型的参数

	var e2 interface{} = studentA // 【写法二】空接口可以接收任意类型的参数 (interface{} 与 T 等价, 因为 T 也是空的)

	fmt.Println(e)
	fmt.Println(e2)
}