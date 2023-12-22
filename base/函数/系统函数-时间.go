package main
import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	nowTime := time.Now()
	fmt.Printf("当前时间为：%v\n", nowTime) // now 返回的是一个结构体


	// 获取年月日
	fmt.Printf("年：%v\n", nowTime.Year())
	fmt.Printf("月：%v\n", nowTime.Month())
	fmt.Printf("日：%v\n", nowTime.Day())
	fmt.Printf("时：%v\n", nowTime.Hour())
	fmt.Printf("分：%v\n", nowTime.Minute())
	fmt.Printf("秒：%v\n", nowTime.Second())


	// 格式化日期时间
	formatTime := nowTime.Format("2006-01-02 15:04:05") //👈 2006-01-02 15:04:05 固定, 不能变!!
	fmt.Printf("当前时间为：%v\n", formatTime) // 2023-08-01 15:04:05


	// 🌟可以根据需求自己任意组合时间的格式, 比如 2006 15:04
	formatTime2 := nowTime.Format("2006 15:04")
	fmt.Printf("当前时间为：%v\n", formatTime2) // 2023 15:04
	


	// 时间常量
	yesterday := nowTime.Add(-time.Hour*24).Day() // 获取一天之前的时间, -time 表示减去时间
	fmt.Printf("一天之前为：%v 号\n", yesterday) 


	// 🔥 得到字符串 (可以的到 当前年份 这几个字以及变量的值！)
	dataStr := fmt.Sprintf("当前年份: %v\n", nowTime.Year())
	fmt.Printf(dataStr)
}