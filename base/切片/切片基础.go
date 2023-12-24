package main
import ("fmt")


/*  
	切片基础:
		1.切片是 Go 特有的概念
		2.切片是建立在数组基础上的抽象, 因为数组的长度不可变, 所以比较不实用, 一般不使用

	切片注意事项:
		1.切片的引用不能越界, 否则会报错
		2.对切片的减少 => 使用切片去切切片 
		3. 🔥如果对生成的切片进行数据的修改, 生成的【新切片】和原来的【数组】、【切片】之间的关系是引用关系, 会统一被修改 !
		4.简写 => [:3]  [3:]  [:]
		
	
*/


func main() {

	fmt.Println("数组转为切片 ————————————————————————————————————————————————")
	// 定义数组 (因为切片构建在数组之上)
	var arr = [7]int{1, 2, 3, 4, 5, 6, 7}


	//  【方式一】从 arr 数组中取出下标为 1 到 5 的元素(包含 1 🌟 不包含 5), 构建一个切片
	// 构建切片
	// var slicesArr[]int = arr[1:5] //

	// 构建切片
	slicesArr := arr[1:5] 

	fmt.Println("切片数据 = ", slicesArr) // [2 3 4 5]
	fmt.Println("切片长度 = ", len(slicesArr))


	// 【方式二】使用 make 函数
	var superSlices = make([]int, 7, 10)
	fmt.Println("superSlices = ", superSlices) // [0 0 0 0 0 0 0]




	fmt.Println("修改切片数据 ————————————————————————————————————————————————")



	// 给切片追加【元素】 (注意, 如果 append 的数据超出了切片的总长, ⚠️ 此时会产生一个新的数组!!)
	slicesArr = append(slicesArr, 88, 99) // append() 函数会返回一个新的切片, 所以需要接收一下
	fmt.Println("切片数据 = ", slicesArr) // [2 3 4 5 88 99]

	// 🔥 给切片追加【切片】
	slicesXXX := arr [1:2] // 2
	slicesArr = append(slicesArr, slicesXXX...) // 把 slicesXXX 切片追加到 slicesArr 切片中
	fmt.Println("追加后的切片数据 = ", slicesArr) // [2 3 4 5 88 99 2] => 2 是 slicesXXX 切片中的元素

	// 🔥 减少切片内的元素 => 使用切片去切切片
	slicesArr = slicesArr[:3] // 从下标为 0 开始 (🔥🔥如果开始下标是 0 , 则可以简写为 [:3], 反之如果结束下标为最尾部, 则可以简写为[3:], 如果是从头到尾, 则可以简写为[:]🔥🔥), 截取到下标为 3 的元素(🌟 不包含!), 也就是截取到 6 为止
	fmt.Println("切片数据 = ", slicesArr) // [2 3 4]

	// 🔥通过切片改变数组元素
	slicesArr[0] = 999
	fmt.Println("切片为: ", arr, "\n", "原先的数组为: ", arr)




	fmt.Println("make 函数定义切片 ————————————————————————————————————————————————")




	// 🚀 使用 make 直接定义定义切片 => 底层是创建了一个数组, 但外界不可见 (访问不到)
	slice := make([]int, 5, 10)  // 🔥 int 类型的切片, 长度为 5, 容量为 10, 如果没有定义值, 则默认为 0
	slice[0] = 111
	slice[1] = 222
	slice[2] = 333
	fmt.Println("slice = ", slice) // slice =  [0 0 0 0 0]




	fmt.Println("遍历切片的数据 ————————————————————————————————————————————————")



	// for 循环遍历切片
	for i := 0; i < len(slice); i ++ {
		fmt.Printf("切片数据: slice[%v] = %v \t \n", i, slice[i]) // \t 表示制表符
	}

	// for range 遍历切片
	for i, v := range slice {
		fmt.Printf("下标: %v, 值: %v \t \n", i, v) // \t 表示制表符
	}



	fmt.Println("复制 A 切片的数据到 B 切片 ————————————————————————————————————————————————")



	slicesA := []int{1, 2, 3, 4, 5} // 👈注意, 如果是空的话 []int 则是一个切片类型!! 如果是 [5]int 则是一个数组类型!!
	slicesB := make([]int, 5, 10)
	copy(slicesB, slicesA) // 把 sliceA 的数据复制到 slicesB 中
	fmt.Println("slicesB = ", slicesB) // slicesB =  [1 2 3 4 5]


}