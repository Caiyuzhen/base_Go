package main
import "fmt"

func main() { 

	// 【🌟 Break 关键字 - 可以结束整个循环】 ————————————————————————————————————————————————————————————————————————————————————————————————
	// 比如累加, 当累加的值超过 300 则暂停
	var sum int = 0
	for i := 1; i <= 100; i ++ {
		sum += 1
		fmt.Println(sum)
		if sum >= 300 {
			break // 停止正在执行的这个循环
		}
	}
	fmt.Println("结束在: ", sum)	
	// 在其他语言内 switch 需要使用 break, 而 go 不用因为 go 会主动帮你加上



	// 【🌟 Break - 只结束离它最近的循环】 双重循环的 break, 循环到内层后, 会等内层循环执行完后再执行外层循环
	for xx := 1; xx <= 5; xx++ {
		for yy := 2; yy <= 4; yy++ {
			fmt.Printf("xx: %v, yy: %v \n", xx, yy)
			if xx == 2 && yy == 2 {
				break // 会把内层的循环停掉而非外层
			}
		}
	}

	fmt.Println("————————————————————————————")


	// 【🌟给循环加标签, 单独停止某个循环标签】
	label2:
	for xx := 1; xx <= 5; xx++ {
		for yy := 2; yy <= 4; yy++ {
			fmt.Printf("xx: %v, yy: %v \n", xx, yy)
			if xx == 2 && yy == 2 {
				break label2 // 会把 label2 这个【外层】循环停掉!!! 
			}
		}
	}

	fmt.Println("————————————————————————————")




	// 【🌟 Continue 关键字 - 结束当前循环并继续下一个循环】 ————————————————————————————————————————————————————————————————————————————————————————————————
	// 比如输出 1～100 以内能被 6 整除的数
	// 【方法一】取余的方式
	for i := 1; i <= 100; i ++ {
		if i % 6 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("————————————————————————————")

	// continue 的方式 => 结束本次循环, 继续下一个循环
	for i := 1; i <= 100; i ++ {
		if i % 6 != 0 {
			continue  // 如果不能整除, 则跳出这个循环, 进行下一个循环 (还会做 i % 6 这件事) => 结束理它近的循环, 开始理它远的循环
		}
		fmt.Println(i)
	}

	fmt.Println("————————————————————————————")

	// 使用 label 结合 container
	label3:
	for i := 1; i <= 100; i ++ {
		if i % 6 != 0 {
			continue label3 
		}
		fmt.Println(i)
	}

	fmt.Println("————————————————————————————")


	// goto 语言 => 无条件的跳转到某个地方 （不建议使用）
	fmt.Println("A")
	if 1 == 1 { // goto 一般都配合条件语句来使用
		goto label4
	}
	fmt.Println("B")
	fmt.Println("C")
	label4:
	fmt.Println("D")


	fmt.Println("————————————————————————————")

	// return 结束当前函数
	for a := 0; a <= 10; a++ {
		if a == 3 {
			fmt.Println("执行到了 a == 3")
			return
		}
	}
	fmt.Println("这句不会执行, 因为整个函数被 return 了")
	
}