package main
import ("fmt")


/* 
	面向对象基础（结构体）:
		1. Go 中没有类, 而是用结构体, 但是 Go 具备面向对象的特效(封装, 继承, 多态)
		2. 结构体是一个【值类型】, 有内存地址

	
	JS 中的类:
		class Person {
			constructor(name = "", age = 0) {
				this.name = name;
				this.age = age;
			}
		}

		const onePerson = new Person("Zeno", 18);


	Python 中的类:
		class Person:
		def __init__(self, name="", age=0):  # 构造函数
			self.name = name
			self.age = age

		one_person = Person("Zeno", 18)
*/


func main() {
	var oneMen Person // 结构体变量但没有传值 【实例化方式一】
	fmt.Println(oneMen) // { 0 } => 初始值为空
	oneMen.name = "Zeno"
	oneMen.age = 18
	fmt.Println(oneMen) // {Zeno 18}


	// 结构体变量 【实例化方式二】
	var teacher = Person{"Kim", 28}
	fmt.Println(teacher) // {Kim 28}


	// 使用【指针】给它指向的值（结构体）进行赋值 【实例化方式三】
	var t *Person = new(Person) // new 表示分配 Person 的内存地址给到 *Person 这个指针 => 🔥🔥 * 是用于【声明】指针类型, 也是用于指针【取值】
	(*t).name = "Annie" // 指针取值
	(*t).age = 30 // 指针取值
	fmt.Println(*t) // {Annie 30}


	// 使用【指针】给它指向的值（结构体）进行赋值 【实例化方式四】
	var x *Person = &Person{}  // 🚀 * 是用于【声明】指针类型, 也是用于指针【取值】 🚀  |🔥🔥 指针接收的一定是一个【内存地址】🔥🔥 & 取地址, 用来获得 Person 的地址, 然后赋值给 x => 直接修改内存中的数据 
	(*x).name = "Hank" // 指针取值
	(*x).age = 34 // 指针取值
}


// 🔥 定义结构体
type Person struct { // 🔥要让外界可以访问的话, 需要首字母大写 !!!
	name string
	age int
}