package main //  main 入口程序
import ( // 🚀导入包一般放在 package 下方 | 且需要用双引号 | 包名一般从 goPath/src 后开始引入 =>  需要配置 GOPATH 环境变量 !! 文件的环境变量!
	"fmt"
	"testPkg/utils" // 需要配置 GOPATH 环境变量 !!
)

func main() {
	// 访问 utils 内的变量
	fmt.Println(utils.StuNo)
}
