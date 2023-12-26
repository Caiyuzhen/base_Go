package model // ⚠️注意包名!!


type Student struct { // 首字母记得大写!!
	Name string
	Age int
}


type animal struct {
	Name string
	Age int
}

// 🌟 使用【工厂模式】可以解决结构体小写外部无法访问的问题
func NewStudent (name string, age int) *animal{ // 🔥返回值为 *animal -> 结构体指针!!! 在方法内部创建结构体实例, 因此外界只要调用 NewStudent 就可以实例化结构体了🔥
	return &animal{name, age} // 🔥 返回结构体的地址 !!!
}