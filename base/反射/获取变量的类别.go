package main

import (
	"fmt"
	"reflect"
)

/* 
	类别:
		类别跟类型不一样, 类别是 reflect.Type.Kind() 或 reflect.Value.Kind() 返回的类型

		🌟 比如类别是 struct, 类型是 main.Cat => 类比【电器】跟【电视机】的关系 🌟
*/

func testReflect(x interface {}) {
	rType := reflect.TypeOf(x)
	rValue := reflect.ValueOf(x)

	// ⚡️ 获取【类别】的方式
	fmt.Println("rType 类别 ", rType.Kind()) // struct
	fmt.Println("rValue 类别 ", rValue.Kind()) // struct


	// ⚠️ 获取类型, 因为 x 在上方已经被转为 interface{} 空接口了, 所以这里需要再转一次 ⚠️
	x2 := rValue.Interface()
	n, flag := x2.(Cat)
	if flag == true {
		fmt.Printf("结构体的类型是:%T ", n) // 🔥 main.Cat
		fmt.Println("猫的名字是: ", n.Name, "年龄是: ", n.Age)
	}


}


type Cat struct {
	Name string
	Age int
}


func main() {
	dudu := Cat {
		Name: "dudu",
		Age: 18,
	}
	testReflect(dudu)
}