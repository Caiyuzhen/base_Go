/* 
	Go 中的运算符
		1. 算术运算符 +, -, *, /, %, ++, --
		2. 赋值运算符 =, +=, -=, *=, /=, %=
		3. 关系运算符 ==, !=, >, <, >=, <=    最终都会返回 true 或 false
		4. 逻辑运算符 &&, ||, !
		5. 位运算符 &, |, ^
		6. 其他运算符 &, *  【🏠 & 可以返回变量的地址】, 【🌟 * 可以返回地址对应的值】
 */
 package main
 import ( "fmt" )

 func main () {
	var n1 int = 4 + 1
	fmt.Println(n1)

	var n2 float64 = float64(n1) / 3
	fmt.Println(n2)

	var n3 int = 10
	var n4 int = 3

	// 取余 a % b 等价于 a - a / b * b
	fmt.Println(n3 % n4) // 3 余 1
	
	// 在 go 里边, ++ 只能单独使用, 不能作为表达式 (不能参与到运算中!) 也不能放在变量前面 !! 
	var n5 int = 6
	n5++
	fmt.Println(n5)

	var num int = 10
	num += 20 // 30 , 等价于 num = num + 20
	fmt.Println(num)

	var num2 int = 100
	num2 /= 10 // 10 , 等价于 num2 = num2 / 10
	fmt.Println(num2)

	var age_1 int = 99
	var age_2 *int = &age_1 // 用 & 把 age_1 的地址赋值给 age_2
	fmt.Println(*age_2) // 通过取地址符 * 可以返回变量的地址

 }