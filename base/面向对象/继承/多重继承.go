package main
import ("fmt")



type A struct {
	a int
	b string
}

type B struct {
	c int
	d string
}

type C struct { // 多重继承
	A
	B
	int // 🔥 结构体的匿名字段可以是基本数据类型 🔥
}







type D struct { // 多重继承
	*A // 👈 也可以嵌入继承结构体的指针
	*B // 👈 也可以嵌入继承结构体的指针
}




type E struct {
	// 👇 表示 x 是 A 类型
	x A // 🌟 结构体的字段可以是结构体类型 | 组合模式
	y B // 🌟 结构体的字段可以是结构体类型 | 组合模式
	z int
}



func main() {
	fmt.Println("结构体的匿名字段可以是基本数据类型 ——————————————————————————————————")


	// 【实例化】方式一:
	cc := C{
		A: A{
			a: 1,
			b: "a",
		},
		B: B{
			c: 2,
			d: "b",
		},
		int: 9999, // 初始化的格式必须要一致!!! 🔥 结构体的匿名字段可以是基本数据类型 🔥
	}

	// 【继承结构体的指针实例化】方式二:
	// cc := C{A{a:1, b:"a"}, B{c:2, d:"b"}, 9999} // 初始化的格式必须要一致!!! 🔥 结构体的匿名字段可以是基本数据类型 🔥

	fmt.Println(cc.A.a)
	fmt.Println(cc.B.c)
	fmt.Println(cc.int)


	fmt.Println("继承结构体的指针实例化 ——————————————————————————————————")


	// 【继承结构体的指针实例化】方式一:
	dd := D{
		A: &A{
			a: 1,
			b: "a",
		},
		B: &B{
			c: 2,
			d: "b",
		},
	}

	// dd := D{&A{1, "a"}, &B{2, "b"}}

	// 【实例化】方式二:


	fmt.Println(*dd.A)
	fmt.Println(*dd.B)
	fmt.Println(dd.a)



	fmt.Println("结构体的字段可以是结构体类型 | 组合模式 ——————————————————————————————————")


	ee := E{ // 组合模式, 不是继承关系
		x: A {
			a: 1,
			b: "a",
		},
		y: B {
			c: 2,
			d: "b",
		},
		z: 9999,
	}

	// fmt.Println(ee.a) // ❌ 因为不是继承关系, 所以无法直接访问 a
	fmt.Println(ee.x.a)
}

