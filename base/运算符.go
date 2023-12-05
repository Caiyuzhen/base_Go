/* 
	Go 中的运算符
		1. 算术运算符 +, -, *, /, %, ++, --
		2. 赋值运算符 =, +=, -=, *=, /=, %=
		3. 关系运算符 ==, !=, >, <, >=, <=
		4. 逻辑运算符 &&, ||, !
		5. 位运算符 &, |, ^
		6. 其他运算符 &, *
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
 }