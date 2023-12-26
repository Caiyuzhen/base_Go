package main
import ("fmt")


func  main() {
	// 创建实例时候直接赋值 (🔥方法一, 缺点则是需要按照顺序)
	var s1 Student = Student{"Jimmy", 22}
	fmt.Println(s1)


	// 创建实例时候直接赋值 (🔥方法二, 按照类型赋值,可以忽略顺序!)
	var s2 Student = Student{
		Age: 22,
		Name: "Kim",
	}
	fmt.Println(s2)

	// 创建实例时候直接赋值 (🚀方法三, 返回结构体指针类型的实例🚀)
	var s3 *Student = &Student{"Annie", 26}
	fmt.Println(*s3)
}


type Student struct {
	Name stringl
	Age int
}