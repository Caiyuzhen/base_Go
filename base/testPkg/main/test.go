package main //  main 入口程序
import ( // 🚀导入包一般放在 package 下方 | 且需要用双引号 | 包名一般从 goPath/src 后开始引入 =>  需要配置 GOPATH 环境变量 !! 文件的环境变量!
	"fmt",
	"base/testPak/utils", // 需要配置 GOPATH 环境变量 !!
)

func main() {
	// 访问 utils 内的变量
	fmt.Println(test.stuNo)
}

对于 Bash（在 macOS 和大多数 Linux 发行版中是默认的），你可以编辑 ~/.bash_profile (macOS) 或 ~/.bashrc (Linux) 文件：

打开终端。
输入 nano ~/.bash_profile (macOS) 或 nano ~/.bashrc (Linux) 来编辑文件。
在文件中添加一行，例如 export GOPATH=/your/go/path。
保存文件并退出编辑器（在 nano 中，使用 Ctrl+O 保存，然后使用 Ctrl+X 退出）。
为了让更改立即生效，运行 source ~/.bash_profile 或 source ~/.bashrc。
对于 Zsh（在最新版的 macOS 中是默认的），你会编辑 ~/.zshrc 文件，步骤类似于 Bash。