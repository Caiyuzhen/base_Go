package main
import "fmt"


type SayHi interface {
	sayHello()
}

type A struct {
	Name string
}

func (a A) sayHello() {
	fmt.Printf("你好, 我是 %v", a.Name)
}

func (a A) fly() { //👈 特有的方法
	fmt.Println("你好, 我会飞")
}

func Greet (p SayHi) { // 👈 多态参数
	p.sayHello()
	p.fly() // 因为这个 fly 是特有的方法, 所以会报错, 🔥 此时需要进行【断言】
}


func main() {
	a := A{"Kim"}
	a.sayHello()
}