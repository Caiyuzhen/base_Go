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
	
	// 🌟 【映射定义方式一】
	var a map[string]string // 定义一个映射

	// 使用 make 进行初始化!! => 应为 var 声明后没有初始化!!
	a = make(map[string]string, 10) // 10 表示可以存放 10 个键值对, 🌟 可以不传, 如果不传的话则会分配 1 个初始空间

	a["name"] = "jimmy"
	a["age"] = "12"
	a["address"] = "NY"

	fmt.Println(a) // address:NY age:12 name:jimmy => 🔥 map 是无序的!!




	// 👍👍 【映射定义方式二】
	b := make(map[int]string, 10)
	b[1] = "学生A"
	b[2] = "学生B"
	b[3] = "学生C"

	fmt.Println(b)



	// 👍👍👍 【映射定义方式三】
	c := map[int]string {
		1: "cat",
		2: "dog",
		3: "rubbit",
	}
	
	fmt.Println(c)




	fmt.Println("—————— 映射的增删改查 ——————")
	fmt.Println("——————————————————————————————————————————————————————————————")




	d := map[int]string {
		1: "catA",
		2: "catB",
		3: "catC",
	}

	// 增
	d[4] = "catD" // map[1:catA, 2:catB, 3:catC, 4:catD]

	// 删 -> 就算没有对应的 key 也不会报错
	delete(d, 3) // 传入 【map 对象】 跟 【key】


	// 修改
	d[1] = "dogA"



	fmt.Println(d) // map[1:dogA, 2:catB, 4:catD]


	// 清空映射 -> make 一个新的 map, 让老的被垃圾回收机制给清空掉
	// newD := make(map[int]string)



	// 查找
	value, bool := d[1] // dogA, true   => 会返回 value 跟 bool, vool 表示是否有这个数据
	value2, _ := d[2] // catB, _  =>  忽略第二个参数
	fmt.Println(value, bool)
	fmt.Println(value2)



	// 获取映射长度
	length := len(d) // 3
	fmt.Println(length)



	// 遍历映射
	for k, v := range d {
		fmt.Printf("key 为: %v value 为: %v \n",k, v)
	}



	fmt.Println("——————————————————————————————————————————————————————————————")



	// 🌟 二维映射的【定义】- 方式一
	aa := make(map[string]map[int]string) // value 也是一个 map
	aa["班级1"] = make(map[int]string)
	aa["班级1"][1] = "Kim"
	aa["班级1"][2] = "Any"
	aa["班级1"][3] = "Zeno"

	aa["班级2"] = make(map[int]string)
	aa["班级2"][1] = "Jimmy"
	aa["班级2"][2] = "Tom"
	aa["班级2"][3] = "Annie"


	// 🌟 二维映射的【定义】- 方式二
	bb := map[string]map[int]string{
		"班级1": {
			1: "Kim",
			2: "Any",
			3: "Zeno",
		},
		"班级2": {
			1: "Jimmy",
			2: "Tom",
			3: "Annie",
		},
	}
	
	fmt.Println(bb)


	// 遍历二维映射
	for k1, v1 := range aa {
		fmt.Println(k1) // k1 为班级1 、 班级2, v1 为 map[int]string

		for k2, v2 := range v1 { // 遍历 v1 , 也就是内层的 map
			fmt.Printf("学号为: %v 姓名为: %v \n", k2, v2) // 学号 姓名
		}
	}

	
}
