package main
import ("fmt")



/*
	🔌 定义接口（定义规范、规则或者某种标准）, 类似 js 中的 interface => 用来规范对象的形状, 但不具体实现
*/
type SayHiInterface interface { // 接口内有的方法, 必须在结构体中实现, 🚀🚀 接口内不能定义变量!!
	sayHello()
}

type BeatInterface interface {
	beat(times int)  // 使 beat 方法接受一个整数参数
}




type Chinese struct {
	Name string
	Age int
}

// 🔥 实现接口的方法
func (men *Chinese) sayHello() { //🚀 在 Go 中, 只要一个类实现了接口的某个方法, 则隐式的被这个接口所约束!! 因此需要实现这个接口内的所有方法, 否则会报错!! 🚀
	fmt.Println("你好，我是中国人")
}




type American struct {
	Name string
	Age int
}

// 🔥 实现接口的方法
func (men *American) sayHello() { // sayHello方法的接收者是一个指向Chinese类型的指针
	fmt.Println("Hi, I'm American")
}




// 🌟 定义专门用来打招呼的函数（非必需, 只是为了统一打招呼）, 接收的是符合 SayeHi 接口标准的变量
func Greet(s SayHiInterface) { // s 是一个接口类型的变量, 可以传入比如 Chinese 或 American 的实例
	s.sayHello()
}




// ——————————————————————————————————————————————————————————————————————————————————————————————————————
// 只要是【🔥自定义数据类型】都可以【实现接口】, 不仅仅是结构体
type Animal int

func (a Animal) sayHello() {
	fmt.Println("Hello, I'm Animal")
}




// ——————————————————————————————————————————————————————————————————————————————————————————————————————
// 🌟 可以实现多个接口
type Person struct {
	Name string
	Age int
}

func (person Person) sayHello() { // 实现一个接口
	fmt.Println("Hello, I'm Person")
}

func (person Person) beat(times int) { // 实现另一个接口 带有参数
	fmt.Printf("I'm beating %v times\n", times)
}



func main() {
	chinese := Chinese{"Lee", 18}
	american := American{"Jack", 20}
	chinese.sayHello()
	american.sayHello()

	Greet(&chinese)
	Greet(&american)




	fmt.Println("——————————————————————————————————————————")


	// 接口本身不可以实例化, 但是可以指向一个实现了该接口的 Class 的实例
	var someone SayHiInterface = &chinese // 因为 chinese 实现了 SayHiInterface 接口, 所以可以赋值给 someone
	someone.sayHello()

	

	fmt.Println("——————————————————————————————————————————")



	// 只要是【🔥自定义数据类型】都可以【实现接口】, 不仅仅是结构体
	var animal Animal = 10
	animal.sayHello()



	fmt.Println("——————————————————————————————————————————")


	person := Person{"Jimmy", 26}
	person.sayHello()
	person.beat(88) // 带有参数的接口
}


