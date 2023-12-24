package main
import ("fmt")

/* 
	二维、多维数组, 本质上都是一维数组, 只是元素类型为一维数组的数组而已
*/

func main()  {
	// 定义二维数组
	var arr[2][3]int // 表示有 2 个一维数组

	// 赋值
	arr[0] = [3]int{1, 2, 3}
	arr[1] = [3]int{4, 5, 6}

	fmt.Println(arr) // [[1 2 3] [4 5 6]]
	


	// 定义二维数组并初始化 - 方式一【手动指定类型】
	var arrX[2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrX) // [[1 2 3] [4 5 6]]

	// 定义二维数组并初始化 - 方式二【自动类型推断】
	var arrZ = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrZ) // [[1 2 3] [4 5 6]]

	// 定义二维数组并初始化 - 方式三【省略长度定义】
	var arrY = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arrY) // [[1 2 3] [4 5 6]]

	// 定义二维数组并初始化 - 方式四【指定索引的值】
	var arrW = [...][3]int{0: {0:999, 1:666, 3}, 1: {4, 5, 6}}
	fmt.Println(arrW) // [[999 666 0] [4 5 6]]


	fmt.Println("------------------------------------------------------------")


	// 二维数组的遍历 - 方式一【普通 for 循环】
	doubleArr := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i < len(doubleArr); i++ {
		for j := 0; j < len(doubleArr[i]); j++ {
			fmt.Printf("doubleArr[%v][%v] = %v\n", i, j, doubleArr[i][j])
		}
	}

	// 二维数组的遍历 - 方式二【for-range 循环】
	for i, v := range doubleArr {
		for j, v2 := range v {
			fmt.Printf("doubleArr[%v][%v] = %v\n", i, j, v2)
		}
	}
}