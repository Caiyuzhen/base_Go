package main
import ("fmt")

/* 
	结构体的方法:
		结构体的方法跟函数是有区别的, 写法也不一样, 方法就是服务于结构体
*/

func main() {
	// 结构体基础 ————————————————————————————————————————————————————————————————
	// var CatA = Cat{"dudu", 3} // 【实例化赋值方法一】
	var CatA = Cat{
		Name: "dudu", 
		Age: 3,
	} // 【实例化赋值方法二】
	CatA.Say() // 喵喵喵

	fmt.Println(CatA) // {dudu 3}  ->   因为 dudu 传入 func(ob Cat)  Say() 时, 是值传递, 所以不会改变原来的值 !!
	fmt.Println(CatA.Name) // 打印 Cat 内的 Name 属性 => dudu (因为是值传递, 所以没有改成 饭团 ~)


	(&CatA).Hi() // 喵喵喵 => 🔥🔥 &CatA 把【内存地址】传入到了 xx 中, 所以这里的 xx 是指针传递, 因此会改变原来的值 !!
	fmt.Println(CatA) // {饭团 3}  ->   因为 dudu 传入 func(xx *Cat)  Hi() 时, 是指针传递, 所以会改变原来的值 !!



	fmt.Println("——————————————————————————————————————————")




	// 绑定在基础数据类型身上的方法 ————————————————————————————————————————————————
	var y2 integer = 20 // 注意 这个 y2 不会被 integer 这个结构体所改变, 因为传入的是【拷贝值!】
	y2.Test2() // 传入 y 的拷贝, 不改变原来的值
	fmt.Println("外部的 y :", y2)

	(&y2).Test3() // 传入 y 的指针， 改变原来的值
	fmt.Println("外部的 y :", y2)




	fmt.Println("——————————————————————————————————————————")




	// 使用 String() 默认打印带有返回值的方法 ————————————————————————————————————————————————————————————————
	var student = Animal{
		Name: "Jimmy",
		Age: 26,
	}
	
	res := &student
	fmt.Println(res)
}


type Cat struct {
	Name string
	Age int
}

// 给结构体绑定方法 (👇里边的 ob 也是定义在这个方法内的)
func(ob Cat) Say() { // ob 为传入的实例化对象(随便取), Cat 为结构体本身 => 🔥🔥 ob 可以是结构体也可以是其他基本数据类型(比如 int, float32 等)
	ob.Name = "饭团" // 🔥🔥 这里的 ob 是值传递, 不会改变原来的值 !!
	fmt.Println("喵喵喵")
}


// 给结构体绑定指针方法（🔥它可以修改原来的值）
func(xx *Cat) Hi() { // *Cat 把传入的参数定义为指针!! 这样就不是值传递了, 🔥🔥 而是【指针传递】了!! 这样就会改变原来的值!!!
	(*xx).Name = "饭团"  // 🔥🔥 这里的 xx 是指针传递, 会改变原来的值 !!
	fmt.Println("喵喵喵")
}





// 把方法绑定在其他数据类型身上 => 不一定要加在结构体身上!也可以加载比如 int 这种数据类型身上!!
type integer int
func(y integer) Test2() { // 传入 y 的拷贝, 不改变原来的值
	y = 99
	fmt.Println("改变后的值（只改变内部的 y）: ", y) // 传入的 y
}

func(y *integer) Test3() { // 传入 y 的指针, 改变原来的值
	*y = 99 // 需要解引用来赋值	!!
	fmt.Println("入 y 的指针（改变实际的 y）: ", *y) // 传入的 y
}







// 🔥 如果绑定的方法为 String(), 【打印实例】则会自动调用此方法!!!
type Animal struct {
	Name string
	Age int
}

func(z *Animal) String() string { // 带有返回值的方法, 返回值为 string  |  🔥 String() 方法会自动调用!!!
	str := fmt.Sprintf("Name = %v, Age = %v", z.Name, z.Age)
	return str
}




