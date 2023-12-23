package main
import ("fmt")

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

	// 🌟 定义一个长度为 5 的数组 =>  // 相当于连续开辟了 5 个 int 类型的空间, 初始值都为 0 
	var scores[5]int
	scores[0] = 95
	scores[1] = 88
	scores[2] = 76
	scores[3] = 66
	scores[4] = 100



	// 求数组的和
	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}


	// 求平均值
	avg := sum / len(scores)

	// 输出结果
	fmt.Printf("总分数为: %v, 平均分为: %v", sum, avg)


	fmt.Println("————————————————————————————————————————————————")


	
	// 🌟 遍历数组
	
}