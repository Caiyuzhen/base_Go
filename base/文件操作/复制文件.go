package main
import (
	"fmt"
	"os"
)

/* 
	🌟 复制文件的逻辑:
		读取文件 => content 存入程序 => 把 content 写入到目标文件夹内
*/

func main() {
	// 要复制的文件夹路径
	filePath := "./file/target.txt"

	// 目标文件夹路径 (复制到哪里)
	targetDirPath := "./file/复制文件的目标文件夹/target.txt"

	//读取文件
	content, err := os.ReadFile(filePath) // 返回【文件内容 []byte】和【错误信息】

	if err != nil {
		fmt.Println("❌ read file err = ", err)
		return
	}

	// 写入文件
	err = os.WriteFile(targetDirPath, content, 0666) // 写入的路径, 写入的内容, 写入的权限
	if err != nil {
		fmt.Println("❌ write file err = ", err)
		return
	} else {
		fmt.Println("✅ 复制文件成功")
	}
}