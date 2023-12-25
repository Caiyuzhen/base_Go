package main
import ("fmt")


func main() {
	var p = Person{"Zeno", 18}
	var s = Student{"Annie", 22}

	// 结构体之间的转换, 需要先转换成同一个类型!!
	s = Student(p) // 把 p 转换成 Student 类型 => 把 p 的值赋值给 s
	fmt.Println(s) // {Zeno 18}


	var p2 = Stu{"Hank", 30}
	fmt.Println(p2) // {Hank 30}
}


type Student struct {
	Name string
	age int
}

type Person struct {
	Name string // 大写是公有的, 外界可以访问
	age int // 🔥 小写是私有的, 外界无法访问
}


type Stu Student // 🔥🔥 把 Student 类型转换成 Stu 类型