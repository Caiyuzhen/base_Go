package main

import (
	"fmt"
	"reflect"
)

/*
	反射(原理类似镜子 🪞, 🌟 反射的的好处是可以传入任意类型的数据 🌟):
		1. 在运行时可以动态获取变量的相关信息, 比如变量的类型信息、类别信息等
		2. 如果是结构体, 则可以获取到结构体本身的信息, 例如结构体的字段、方法等
		3. 通过反射, 可以修改变量的值, 可以调用关联的方法
*/

// 🔥 把函数的参数定义为【空接口】, 可以接收任意类型的值
func testReflect(i interface{}) {
	// A: 调用 TypeOf 函数, 返回 reflect.Type 类型的数据
	rType := reflect.TypeOf(i)
	fmt.Println("类型 = ", rType) // rType =  int, 是 TypeOf(i) 类型, ⚠️ 不能直接使用 !!
	//如果要获取数值

	// B : 调用 ValueOf 函数, 返回 reflect.Value 类型的数据
	rValue := reflect.ValueOf(i).Int() // 🔥 【让反射的 value 可用的方法一】Int() 方法返回 int64 类型的数据!!
	fmt.Println("值 = ", rValue)        // 值 =  999, 是 ValueOf(i) 类型, ⚠️ 不能直接使用 !! 要使用得调用 Int() 方法!!!

	res := rValue + 1
	fmt.Println("res = ", res) // 1000




	// 🔥【让反射的 value 可用的方法二】先将Value 转为空接口
	i2 := rValue.Interface()

	// 将空接口断言为 string 类型, 用 x 接收
	x := i2.(string) // 👈 类型断言, 断言为 string 类型
	fmt.Println("x 的值 ", x)

}

func main() {
	// 对基本数据类型进行反射
	num := 999
	testReflect(num)

}
