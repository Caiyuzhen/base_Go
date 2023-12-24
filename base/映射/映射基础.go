package main
import ("fmt")

/* 
	注意:
		1.🌟 Go 的映射要求所有的 key 和 value 都是相同的类型。这由语言的静态类型特性所强制
		2.🌟 映射中的元素是无序的。
		3. key 跟 value 的类型可以是 bool 、str 、数字、指针、通道、接口、结构体、数组 等 ( 🔥 但是 key 不可以是 切片、映射、函数 )
		4. 只声明 map 是不会分配内存空间的, 🔥 需要使用 make 来【初始化】分配内存空间
		5. key 不能重复, 如果重复了, 后一个则会覆盖前一个


	Python 中的字典
		# 定义字典
		capitals = {
			"France": "Paris",
			"Italy": "Rome",
			"Japan": "Tokyo"
		}

		# 添加键值对
		capitals["Germany"] = "Berlin"


	JS 中的对象
		// 定义对象
		const capitals = {
			France: "Paris",
			Italy: "Rome",
			Japan: "Tokyo"
		};

		// 添加键值对
		capitals["Germany"] = "Berlin";
*/

func main() {

	fmt.Println("—————— 映射的定义 ——————")
	// 定义方式
	var a map[string]string // 定义一个映射

	// 🌟 使用 make 进行初始化
	a = make(map[string]string, 10) // 10 表示可以存放 10 个键值对, 🌟 可以不传, 如果不传的话则会分配 1 个初始空间

	a["name"] = "jimmy"
	a["age"] = "12"
	a["address"] = "NY"

	fmt.Println(a) // address:NY age:12 name:jimmy => 🔥 map 是无序的!!
	
}