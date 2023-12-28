package main
import (
	"fmt"
	"os"
)


func main() {
	// 打开文件
	file, err := os.Open("./file/text.txt") // 返回【文件指针】和【错误信息】
	if err != nil { // 如果不等于空, 说明有错误
		fmt.Println("open file err = ", err)
	} else {
		fmt.Println("file = ", file) // 🔥 打印的是【文件指针】

		err2 := file.Close()// 关闭文件
		if err != nil { // 如果不等于空, 说明有错误
			fmt.Println("close file err = ", err2)
		}
	}
}