package main
import (
	"fmt"
	"os"
)

/* 
	IO 流
		流:
			text.txt => 程序 => 目标文件

		- 输入流: 从外部读取数据到程序中  
		- 输出流: 从程序输出数据到外部


	Next JS 读取文件的 api:
		// pages/api/readfile.js
		import fs from 'fs';
		import path from 'path';

		export default function handler(req, res) {
			const filePath = path.join(process.cwd(), 'example.txt');
			fs.readFile(filePath, 'utf8', (err, data) => {
				if (err) {
				res.status(500).json({ message: 'Error reading file' });
				return;
				}
				res.status(200).json({ content: data });
			});
		}

	Node JS 读取文件:
		const fs = require('fs');

		fs.readFile('example.txt', 'utf8', (err, data) => {
		if (err) {
			console.error('Error reading the file:', err);
			return;
		}
		console.log(data);
		});

	Python 读取文件:
		try:
			with open('example.txt', 'r') as file:
				content = file.read()
				print(content)
		except FileNotFoundError:
			print("The file was not found")
*/

func main() { // 🔥🔥 使用 os.ReadFile 来读取文件不需要用 open() 跟 close(), 因为 ReadFile 函数已经做了
	// 读取文件
	content, err := os.ReadFile("./file/text.txt") // 返回【文件内容 []byte】和【错误信息】
	if err != nil { // 如果返回的不是 0 => 不是 true
		fmt.Println("read file err = ", err)
	} else {
		// 把内容显示在终端
		fmt.Printf("content = %v", string(content)) // 🔥 因为返回的是 []byte 的切片, 所以要用 string 方法转为1字符串 => 打印的是【文件内容】
	}
}