package main
import (
	"fmt"
	"unsafe"
	"strconv"
)

func main() {
	// 有符号的【整型】
	var num1 int8 = 127 // 1 字节 (-128~127)
	var num2 int16 = -32767// 2 字节 (-32768~32767~32768)
	var num3 int32 =  -214783647// 4 字节 (-214783648~214783647~214783648~214783647)
	var num4 int64 = -9999999999999999 // 8 字节 (-2^63~2^63-1)

	// 无符号的【整型】=> 比如年龄, 不会有负的, 也不会超出 200 岁
	var num5 uint8 = 255// 1 字节 (0~255)
	var num6 uint16 = 22222 // 2 字节 (0~2^16-1)
	var num7 uint32 = 9999999 // 4 字节 (0~2^31-1)
	var num8 uint64 = 99999999999 // 8 字节 (0~2^63-1)

	fmt.Println(num1, num2, num3, num4, num5, num6, num7, num8)

	fmt.Printf("num4 的类型是: %T", num4) // 🔥 Printf 用来格式化字符串 | 🔥 %T 是用来填充 num4 的占位, 表示类型


	// 其他类型
	var num9 int // 等价于 int64 或 int32 => 取决于操作系统
	var num10 uint // 等价于 uint64 或 uint32
	var num11 rune // 等价于 int32
	var num12 byte // 等价于unit8 (0~255)
	fmt.Println()
	fmt.Println(num10, num11, num12)
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num9)) //🔥查看类型



	// 类型转换 ——————————————————————————————————————————————————————————————————————————————————————————
	// 基本类型（int、float、bool 等）转换为 str 类型
	// int -> str
	var n1 int = 19
	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("s1 的类型是: %T , s1 = %q \n", s1, s1)

	// float -> str
	var n2 float64 = 4.29
	var s2 string = strconv.FormatFloat(n2, 'f', 9, 64) // ⚡️ f 表示转为十进制, 9 表示保留小数后 9 位, 64 表示 loar64 类型
	fmt.Printf("s1 的类型是: %T , s2 = %q \n", s2, s2)

	// bool -> str
	var n3 bool = true
	var s3 string = strconv.FormatBool(n3)
	fmt.Printf("s3 的类型是: %T , s2 = %q \n", s3, s3)




	// str 转为基本类型（int、float、bool 等）
	// str -> bool
	var s11 string = "true"
	var b11 bool
	b11, _ = strconv.ParseBool(s11) // 🔥 因为 ParseBool 的返回值有两个 => (value bool, err error), 所以用 b 跟 _ 来接收返回值
	fmt.Printf("b11 的类型是: %T, b11 = %v \n", b11, b11)

	// str -> int64
	var s22 string = "19"
	var num11 int64
	num11 = strconv.parseInt(s22, 10, 64) // 10 表示十进制, 64 表示 int64
	fmt.Printf("num11 的类型是: %T, num11 = %v \n", num11, num11)

	// str -> float32 / float64
	var s33 strint = "3.14"
	var num22 float64
	f1, _ = strconv.parseFloat(s33, 64) // 转为 float64 位, 🔥 因为 ParseFloat 的返回值有两个 => (value bool, err error), 所以用 f1 跟 _ 来接收返回值

}