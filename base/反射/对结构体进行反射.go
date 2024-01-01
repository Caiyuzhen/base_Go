package main

import (
	"fmt"
	"reflect"
)

/* 
	🌟 反射的的好处是可以传入任意类型的数据 🌟
		比如把某个结构体实例反射给 x, 用 x 接收
*/


// 🔥 把函数的参数定义为【空接口】, 可以接收任意类型的值
func testReflect(i interface{}) {
	rValue := reflect.ValueOf(i)
	fmt.Println("值 = ", rValue) 


	// 🔥【让反射的 value 可用的方法二】先将Value 转为空接口
	i2 := rValue.Interface() // ⚠️ 因为 i 在上方已经被转为 interface{} 空接口了, 所以这里需要再转一次 ⚠️

	// 将空接口断言为 string 类型, 用 x 接收
	x, flag := i2.(Student) // 👈 类型断言, 断言为 Student 类型, 因为不一定是 Studeng 类型, ⚠️ 所以用 flag 接收是否断言成功
	if flag == true {
		fmt.Printf("学生的名字是: %v \n", x.Name)
		fmt.Printf("学生的年龄是: %v \n", x.Age)
	}
}


type Student struct {
	Name string
	Age int
}


func main() {
	// 对结构体进行反射
	stuentA := Student {
		Name: "小明",
		Age: 18,
	}
	testReflect(stuentA)

}
