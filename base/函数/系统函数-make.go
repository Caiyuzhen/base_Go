package main
import ("fmt")


/* 
	使用 make 函数来分配内存, 定义切片
*/

func main() {
	// 定义切片（使用 make 直接定义 => 底层是创建了一个数组, 但外界不可见 (访问不到))
	slice := make([]int, 5, 10)  // int 类型的切片, 长度为 5, 容量为 10 🔥
	fmt.Println("slice = ", slice) // slice =  [0 0 0 0 0]
	fmt.Println("slice 的长度 = ", len(slice)) // 5
	fmt.Println("slice 的容量 = ", cap(slice)) // 10

	slice[0] = 99
	slice[1] = 88

	// 🔥 给切片追加元素 (注意, 如果 append 的数据超出了切片的总长, ⚠️ 此时会产生一个新的数组!!)
	slice = append(slice, 77, 66, 55) // 🔥 追加是向后追加!!
	fmt.Println("slice = ", slice) // [99 88 0 0 0 77 66 55]
}