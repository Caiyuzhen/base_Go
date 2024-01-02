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

// 通过反射修改【结构体】的值 Field() 方法跟 NumMethod() 方法
func changeStructReflect(x interface{}) {
	reValue := reflect.ValueOf(x) // x 为传入的结构体

	// 👇 调用反射结构体内的方法 ————————————————————————————————————————————————————————————————————————
	fmt.Println("👇 调用反射结构体内的方法 ————————————————————————————————————————————————————————————————————————")
	field := reValue.Elem().NumField()  // 👈 Elem() 可以获取引用, 🌟 NumField() 可以获取结构体有多少个字段数量
	for i := 0; i < field; i++ {
		fmt.Printf("第 %d 个字段的值是: %v\n", i, reValue.Elem().Field(i))
	}

	methodNum := reValue.NumMethod() //👈 获取结构体的方法数量
	for i := 0; i < methodNum; i++ {
		fmt.Printf("第 %d 个方法的名称是: %v\n", i, reValue.Method(i).Type())
	}

	reValue.Method(1).Call(nil) // 调用结构体内的方法 => 第一个方法 => 要调用方法的话, 方法的首字母必须大写!! 🔥 同时方法的顺序按照 ASCII 码的顺序进行排列(a, b, c, d, ...)

	var params []reflect.Value	
	params = append(params, reflect.ValueOf(10)) // 因为是 []reflect.Value 类型, 所以需要 reflect.Valueof(10) 来传参
	params = append(params, reflect.ValueOf(20))
	result := reValue.Method(0).Call(params) // 👈 由于 Call 方法要求传入一个 value 切片, 因此需要通过上面的步骤定义切片并传入数据, 返回值 result 也是一个切片
	fmt.Printf("调用方法后获得的返回值(切片)为: %v", result[0].Int()) // 🔥 reflect 的返回值, 需要调用 Int() !!




	// 👇 修改反射结构体内的属性 ————————————————————————————————————————————————————————————————————————
	fmt.Println("👇 修改反射结构体内的属性 ————————————————————————————————————————————————————————————————————————")
	// SetInt(22)  | SetString("Jimmy") | ...
	z := reValue.Elem().NumField() // 2 个属性
	fmt.Printf("该结构体内共有 %v 个属性 \n", z)
	reValue.Elem().Field(1).SetInt(22) // 修改结构体内的年龄, 👉 顺序是按照结构体定义的顺序 !!
	reValue.Elem().Field(0).SetString("饭团")
}




// 结构体
type Cat struct {
	Name string
	Age int
}

// 给结构体绑定方法
func (cat Cat) GetSum(n1, n2 int) int { // 返回值是 int
	fmt.Println("调用了 GetSum 方法")
	return n1 + n2
}

func (cat Cat) Say() {
	fmt.Println("猫喵叫")
	fmt.Println("猫的名字是: ", cat.Name, "年龄是: \n", cat.Age)
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
	fmt.Printf("名称修改后的值: %v \n", dudu.Name)
	fmt.Printf("年龄修改后的值: %v \n", dudu.Age)
}