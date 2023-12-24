package main
import ("fmt")

/* 
	数组的基础知识:
		1.数组的长度是类型的一部分, 因此具有不同长度的数组为不同类型
		2.🔥 在 Go 中, 数组是【值类型】=> 而不是【引用类型】!! 跟 JS 还有 Python 不一样!! 
		3.🔥 如果要改变原数组的值, 需要使用【指针】来引用传递 => 传递内存地址!!
		4.🔥 不能移除数据, 只能改变数据, 因为数组的长度是固定的 !!
*/


/* 
	TS 的泛型数组:
		function getMax<T>(arr: T[]): T {
			let max = arr[0];
			for (let i = 1; i < arr.length; i++) {
				if (arr[i] > max) { // 这里假设了T支持比较操作
					max = arr[i];
				}
			}
			return max;
		}

		// 使用时指定T为具体的类型
		let maxNumber = getMax<number>([1, 3, 2]);
		let maxString = getMax<string>(['a', 'b', 'c']);
*/

func main() {

	// 🌟 【【声明且初始化方式一 - 手动指定类型】
	var superArr[3]int = [3]int{1, 2, 3}
	fmt.Println("superArr = ", superArr) // superArr =  [1 2 3]
	

	// 🌟 【声明且初始化方式二 - 自动类型推断】
	var superArr2 = [3]int{1, 2, 3}
	fmt.Println("superArr2 = ", superArr2) // superArr2 =  [1 2 3]


	// 🌟 【声明且初始化方式三 - 省略长度定义】
	var superArr3 = [...]int{1, 2, 3}
	fmt.Println("superArr3 = ", superArr3) // superArr3 =  [1 2 3]


	// 🌟 【声明且初始化方式四 - 指定索引的值】
	var superArr4 = [...]int{0: 888, 1: 996, 2: 777}
	fmt.Println("superArr4 = ", superArr4) // superArr4 =  [888 996 777]


	// 👇 同理, 定义的时候可以使用 :=
	superArr5 := [...]int{1, 2, 3}
	fmt.Println("superArr5 = ", superArr5) // superArr5 =  [1 2 3]



	fmt.Println("————————————————————————————————————————————————")
	// 改变数组内的某个数据 (🔥不能移除数据, 只能改变数据, 因为数组的长度是固定的 !!)
	superArr4[0] = 666





	fmt.Println("————————————————————————————————————————————————")


	// 🌟 定义一个长度为 5 的数组 =>  // 相当于连续开辟了 5 个 int 类型的空间, 初始值都为 0 
	var scores[5]int
	scores[0] = 95
	scores[1] = 88
	scores[2] = 76
	scores[3] = 66
	scores[4] = 100


	// 遍历并求数组的和
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}


	// 求平均值
	avg := sum / len(scores)

	// 输出结果
	fmt.Printf("总分数为: %v, 平均分为: %v \n", sum, avg)



	fmt.Println("————————————————————————————————————————————————")


	
	// 🌟 从键盘录入数据到数组内
	var arrs[5]int
	result := 0

	for i := 0; i < len(arrs); i ++ {
		fmt.Println("请输入第", i + 1, "个元素的值: ")
		fmt.Scanln(&arrs[i]) // 👈 从【键盘录入】数据到数组内, 🔥 通过 &arrs[i] 来获取数组元素的内存地址
	}
	
	// 【遍历方式一】for 循环遍历数组
	for i := 0; i < len(arrs); i ++ {
		result += arrs[i]
	}

	fmt.Println("数组的和为: ", result, "平均值为: ", result / len(arrs))



	fmt.Println("————————————————————————————————————————————————")


	// 【遍历方式二-1】 用 for range 来遍历数组
	for key,value := range arrs { // key 跟 value 可以随便写, 但是必须都写上
		fmt.Printf("arrs[%v] = %v \n", key, value)
	}


	fmt.Println("————————————————————————————————————————————————")

	// 【遍历方式二-2】 忽略 key, 只打印 value 的方式
	for _,value := range arrs { // key 跟 value 可以随便写, 但是必须都写上
		fmt.Printf("成绩为: = %v \n", value)
	}



	fmt.Println("数组为值类型 ————————————————————————————————————————————————")
	
	// 🔥 在 Go 中, 数组是【值类型】=> 而不是【引用类型】!!
	arrChange := [3]int{1, 2, 3}
	test(arrChange)
	fmt.Println("arrChange 的值未改变 = ", arrChange) // arrChange =  [1 2 3] => 🔥说明数组是值类型, 传递的是值拷贝, 不会改变原数组的值


	fmt.Println("使用指针来改变引用的数组 ————————————————————————————————————————————————")
	arrChange2 := [3]int{1, 2, 3}
	test2(&arrChange2) // 传递内存地址
	fmt.Println("arrChange2 的值改变了 = ", arrChange2) // arrChange2 =  [88 2 3] => 🔥说明数组是值类型, 传递的是值拷贝, 不会改变原数组的值

}



// 🔥 在 Go 中, 数组是【值类型】=> 而不是【引用类型】!!
func test(arr [3]int) { // 会在函数的栈内存中开辟一块新的内存空间, 用来存储 arr 的值拷贝
	arr[0] = 88 // 改变的是拷贝的值!
}

func test2(arr *[3]int) { // 传递的是指针, 会改变原数组的值
	(*arr)[0] = 88 // 改变的是引用的值!
}