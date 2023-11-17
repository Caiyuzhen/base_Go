package main
import (
	"fmt"
	"unsafe"
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
}