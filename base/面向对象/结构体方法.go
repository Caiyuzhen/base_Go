package main
import ("fmt")

/* 
	结构体的方法:
		结构体的方法跟函数是有区别的, 写法也不一样, 方法就是服务于结构体
*/

func main() {
	var CatA = Cat{"dudu", 3}
	CatA.Say() // 喵喵喵

	fmt.Println(CatA) // {dudu 3}  ->   因为 dudu 传入 func(ob Cat)  Say() 时, 是值传递, 所以不会改变原来的值 !!
	fmt.Println(CatA.Name) // 打印 Cat 内的 Name 属性 => dudu (因为是值传递, 所以没有改成 饭团 ~)


	(&CatA).Hi() // 喵喵喵 => 🔥🔥 &CatA 把【内存地址】传入到了 xx 中, 所以这里的 xx 是指针传递, 因此会改变原来的值 !!
	fmt.Println(CatA) // {饭团 3}  ->   因为 dudu 传入 func(xx *Cat)  Hi() 时, 是指针传递, 所以会改变原来的值 !!
}


type Cat struct {
	Name string
	Age int
}

// 给结构体绑定方法 (👇里边的 ob 也是定义在这个方法内的)
func(ob Cat) Say() { // ob 为传入的实例化对象(随便取), Cat 为结构体本身
	ob.Name = "饭团" // 🔥🔥 这里的 ob 是值传递, 不会改变原来的值 !!
	fmt.Println("喵喵喵")
}


// 给结构体绑定指针方法（🔥它可以修改原来的值）
func(xx *Cat) Hi() { // *Cat 把传入的参数定义为指针!! 这样就不是值传递了, 🔥🔥 而是【指针传递】了!! 这样就会改变原来的值!!!
	xx.Name = "饭团"  // 🔥🔥 这里的 xx 是指针传递, 会改变原来的值 !!
	fmt.Println("喵喵喵")
}
