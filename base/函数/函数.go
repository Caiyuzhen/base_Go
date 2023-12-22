package main
import "fmt"

// 🌟 在 Go 中【基本数据类型】跟【数组】都是进行【值传递】=> 拷贝传递而非引用传递, 所以不会改变原来的值, 这点跟 JS 还有 Python 不同
// 🔥 要实现相同的能力, 则需要传入指针 => 比如数组的指针，允许函数修改原数组的内容
/* 
	GO:
	func modifyArray(arr *[3]int) {
		arr[2] = 100  // 修改原数组的第三个元素
	}

	func main() {
		originalArray := [3]int{1, 2, 3}
		modifyArray(&originalArray)
		fmt.Println(originalArray)  // 输出: [1, 2, 100]
	}

	Python:
		def modify_list(lst):
		lst.append(4)  # 修改原始列表

		original_list = [1, 2, 3]
		modify_list(original_list)
		print(original_list)  # 输出: [1, 2, 3, 4]

	JS:
		function modifyArray(arr) {
			arr.push(4);  // 修改原始数组
		}
		
		let originalArray = [1, 2, 3];
		modifyArray(originalArray);
		console.log(originalArray);  // 输出: [1, 2, 3, 4]		
*/

func main() {
	num1, num2 := 10, 20
	res := calculate(num1, num2)
	fmt.Println(res)

	resA, resB := calculate2(num1, num2)
	fmt.Println(resA, resB)

	// 🔥 可以用 _ 忽略其中一个返回值, 只接收一个值, 忽略另一个值
	resC, _ := calculate2(num1, num2)
	fmt.Println(resC)



	// 传递多个值
	multiValue(1, 2, 3, 4)


	// 🌟 Case3: 改变原始数组的值 => 传入指针的地址
	numOld := 999
	fmt.Println("numOld 的初始值为:", numOld) // 999
	changeValue(&numOld) // 👈 传入 num3 的指针地址, 改变原始值!
	fmt.Println("numOld 的新值为:", numOld) // 输出: 300


	// 🌟 Case4: 函数也是一种数据类型, 可以赋值给一个变量
	a := TestCase
	a(19) // 等价于调用了 TestCase(19)



	// 🌟 Case5: 传入函数行参
	TestCase2(11, 12, a)


	// 🌟 Case6: 使用类型别名定义的函数
	TestCase3(11, TestCase)


	// 🌟 Case7: 自动识别返回值（不用关注返回值的顺序）
	resAA, resBB := superTest(10.5, 20) // 输出: 9.5  30.5
	fmt.Println("resAA :", resAA, "resBB :", resBB)
}





// Case 1: 返回值 ————————————————————————————————————————————————————————————————————————————————————————————————
// 计算和的函数 => 返回一个值
func calculate(num1 int, num2 int) int {// (int) 为返回值类型, 如果返回值类型就一个的话, 可以省略 () 括号
	var res int = 0
	res += num1
	res += num2
	return res
}


// 可以计算和也可以计算差的函数 => 返回多个值
func calculate2(num1 int, num2 int) (int, int) {
	var resA int = 0
	var resB int = 0
	resA += num1
	resA += num2

	resB = num1 - num2

	return resA, resB
}


// 🌟 Case2: 支持可变参数（参数的数量可变 => 使用 ...）———————————————————————————————————————————————————————————————————————————————————————————————
func multiValue(args...int) { // 👈 ...表示可以传入任意数量的可变参数
	fmt.Println("--------------------------------")
	// 处理可变参数 => 把可变参数当成切片处理
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}


// 🌟 Case3: 改变原始数组的值 => 传入指针的地址 ———————————————————————————————————————————————————————————————————————————————————————————————
func changeValue(num *int) {
	*num = 300 // 🔥🔥使用 *num 来修改指针指向的值
}


// 🌟 Case4: 函数也是一种数据类型, 可以赋值给一个变量 ———————————————————————————————————————————————————————————————————————————————————————————————
func TestCase(num int) {
	fmt.Println("TestCase 函数被调用了, 获得了值: ", num)
}


// 🌟 Case5: 🚀 可以把函数作为一个行参来传入
func TestCase2(num1 int, num2 float32, TestCase func(int)) { // TestCase func(int) 表示传入函数
	fmt.Println("————————————————————————————————————————")
	fmt.Println("TestCase2 获得了", num1, num2)
	TestCase(num1) // 调用函数！传入 num1 到 TestCase 函数中！
}

/* 
	Python:
		from typing import Callable // 在Python中，Callable类型用于类型注解，用来指定一个函数作为参数, Callable[[Arg1Type, Arg2Type, ...], ReturnType]的格式用于定义一个函数类型，其中Arg1Type, Arg2Type, ...表示该函数接受的参数类型，而ReturnType则表示该函数的返回类型

		def test_case(num: int) -> None:
			print("test_case 函数被调用了，获得了值:", num)

		def test_case2(num1: int, num2: float, func: Callable[[int], None]) -> None:
			print("————————————————————————————————————————")
			print("获得了", num1, num2)
			func(num1)  # 调用传入的函数
	test_case2(10, 20.5, test_case)

	TS:
		function testCase(num: number): void {
			console.log("testCase 函数被调用了，获得了值:", num);
		}

		function testCase2(num1: number, num2: number, func: (n: number) => void): void {
			console.log("————————————————————————————————————————");
			console.log("获得了", num1, num2);
			func(num1); // 调用传入的函数
		}

		testCase2(10, 20, testCase);
*/

// 🌟 Case6: 自定义数据类型 (类似类型别名) ———————————————————————————————————————————————————————————————————————————————————————————————
type myInt int // 自定义变量类型 myInt
// var abc myInt = 39
// fmt.Println(abc)

type myFunc func(int)  // 自定义函数类型 myFunc, 可以简化如下定义
func TestCase3(num1 int, TestCase myFunc) { 
	fmt.Println("---------")
	fmt.Println("TestCase3 获得了", num1) // 调用函数！传入 num1 到 TestCase 函数中！
}

/* 
	Python:
		# 定义一个函数类型 myFunc，它接受一个整数参数并且没有返回值
		myFunc = Callable[[int], None]

		# 使用 myFunc 类型的函数作为参数
		def test_case3(num1: int, test_case: myFunc) -> None:
			print("TestCase3 获得了", num1)
			test_case(num1)
	TS:
		// 定义一个新的类型 myInt
		type myInt = number;

		// 定义一个函数类型 myFunc，它接受一个整数参数并且没有返回值
		type myFunc = (num: number) => void;

		// 使用 myFunc 类型的函数作为参数
		function testCase3(num1: myInt, testCase: myFunc): void {
			console.log("TestCase3 获得了", num1);
			testCase(num1);
		}
*/

 
// 🌟 Case7: 自动识别返回值（不用关注返回值的顺序） ———————————————————————————————————————————————————————————————————————————————————————————————
func superTest(num1 float32, num2 int) (resA float32, resB float32) {
	resA = float32(num2) - num1 // 需要转换类型才能用【整型】来减去【浮点数】！
	resB = num1 + float32(num2) // 需要转换类型才能用【整型】加上【浮点数】！
	return // 👈 这种写法就不用写 return resA, resB , 且不用跟上面的类型定义一一对应其位置
}
