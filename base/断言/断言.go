package main
import "fmt"


type SayHi interface {
	sayHello()
}

type A struct {
	Name string
}

type B struct {
	Name string
}


func (a A) sayHello() {
	fmt.Printf("你好, 我是 %v \n", a.Name)
}

func (a A) fly() { //👈 特有的方法
	fmt.Println("你好, 我会飞")
}

func (b B) sayHello() {
	fmt.Printf("你好, 我是 %v \n", b.Name)
}

func (b B) jump() { //👈 特有的方法
	fmt.Println("你好, 我会跳")
}



func Greet (person SayHi) { // 👈 多态参数, 继承 SayHi 接口
	person.sayHello()

	// 进行【断言】方法一, 不做判断, 可能会报错: 
	// var ch A = person.(A) // 👈👈 让 p 转为 A 类型, 然后赋值给 ch 去执行 fly() 方法 => 用来判断 p 是否是 A 类型
	// ch.fly() // 因为这个 fly 是特有的方法, 而 B 没有, 所以会报错, 🔥 此时需要进行【断言】, p判断是否是 A 类型, 是 A 类型才执行 fly() 方法



	// 进行【断言】方法二, 做判断, 更健壮:
	ch2, ok := person.(A) // 👈👈 让 p 转为 A 类型, 然后赋值给 ch 去执行 fly() 方法 => 用来判断 p 是否是 A 类型
	if ok { // ok 是一个布尔类型  
		ch2.fly()
	} else {
		fmt.Println("不是 A 类型, 无法执行 fly() 方法")
	}



	// 进行【断言】方法三, 更简略的 Go 写法:
	 // 👈👈 让 p 转为 A 类型, 然后赋值给 ch 去执行 fly() 方法 => 用来判断 p 是否是 A 类型
	if ch3, ok := person.(A); ok { // ok 是一个布尔类型  
		ch3.fly()
	} else {
		fmt.Println("不是 A 类型, 无法执行 fly() 方法")
	}


	// 🌟 进行多重判断（可以断言多个类型）:
	switch person.(type) { // 🔥🔥 type 是 go 的关键字, 固定写法, 用来判断 ch4 的类型
		case A:
			ch4 := person.(A)
			ch4.fly()
		case B:
			ch4 := person.(B)
			ch4.jump()
	}
}


func main() {
	a := A{"Kim"}
	b := B{"Zeno"}
	// a.sayHello()
	Greet(a)
	Greet(b)
}