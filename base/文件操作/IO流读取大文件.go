package main
import (
	"fmt"
	"os"
	"io"
	"bufio"
)



/* 
	🔥 读取大文件

	Python 的文件迭代器读取大文件:
		with open('largefile.txt', 'r') as file:
		for line in file:
			# 处理每一行
			print(line)

	Node 的 Stream 读取大文件:
		const fs = require('fs');
		const readline = require('readline');

		const stream = fs.createReadStream('largefile.txt');
		const reader = readline.createInterface({ input: stream });

		reader.on('line', (line) => {
		console.log(line);
		});

		reader.on('close', () => {
		console.log('Finished reading the file');
		});

*/

func main() {
	// 打开文件
	file, err := os.Open("./file/大文件.txt")
	if err != nil { // 如果返回的不是 0 => 不是 true
		fmt.Println("打开失败")
	} 
		
	// 当函数退出时, 关闭文件, 放置内存泄漏
	defer file.Close()

	// 🌟创建一个缓冲区(流)，把文件内容放入缓冲区中
	reader := bufio.NewReader(file)

	// ♻️ 循环读取
	for {
		str,err := reader.ReadString('\n') // 🔥 => 这里的 \n 是字符(换行符), 所以需要使用单引号!!!  '\n' 表示以换行作为一次读取, 到换行则结束一次读取
		if err == io.EOF { // 👈 读取到结尾
			break
		}
		// 🔥 如果没读到末尾, 则正常输出内容
		fmt.Println(str)
	}

	// 结束
	fmt.Println("✅ 文件读取完毕")
}