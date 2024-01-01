package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改基本数据类型的反射
func changeBaseNumDataReflect(x interface{}) {
	reValue := reflect.ValueOf(x)
	reValue.Elem().SetInt(200) // 👈 Elem() 获取指针的值!! 把引用的值改为 200
}

// 通过反射修改结构体的值
func changeStructReflect(x interface{}) {
	reValue := reflect.ValueOf(x)

	field := reValue.Elem().NumField()  // 👈 获取结构体有多少个字段数量
	for i := 0; i < field; i++ {
		fmt.Printf("第 %d 个字段的值是: %v\n", i, reValue.Elem().Field(i))
	}

	methodNum := reValue.NumMethod() //👈 获取结构体的方法数量
	for i := 0; i < methodNum; i++ {
		fmt.Printf("第 %d 个方法的名称是: %v\n", i, reValue.Method(i).Type())
	}

	reValue.Method(1).Call(nil) // 调用结构体内的方法 => 第一个方法 => 要调用方法的话， 方法的首字母必须大写!!
}




// 结构体
type Cat struct {
	Name string
	Age int
}

// 给结构体绑定方法
func (cat Cat) Say() {
	fmt.Println("猫喵叫")
	fmt.Println("猫的名字是: ", cat.Name, "年龄是: \n", cat.Age)
}

func (cat Cat) GetSum(n1, n2 int) int { // 返回值是 int
	return n1 + n2
}

func (cat Cat) Set(name string, age int) { // 设置结构体实例名称的方法
	cat.Name = name
	cat.Age = age
}




func main() {
	var num int = 100
	changeBaseNumDataReflect(&num) // 🔥传入指针, 修改的是引用地址的值
	fmt.Printf("反射修改后的值: %v \n", num)


	dudu := Cat {
		Name: "dudu",
		Age: 18,
	}
	changeStructReflect(&dudu)
}