package main
import "fmt"
/* 
	🌟 多态指实例具有多种形态
			Go 中的多态适应接口来呈现


	TS 中多态的例子:
		// 定义一个抽象类
		abstract class Animal {
			// 定义一个抽象方法
			abstract makeSound(): void;
		}

		// 定义两个子类，继承自Animal
		class Dog extends Animal {
			makeSound() {
				console.log("Woof! Woof!");
			}
		}

		class Cat extends Animal {
			makeSound() {
				console.log("Meow! Meow!");
			}
		}

		// 实例化子类
		let dog: Animal = new Dog();
		let cat: Animal = new Cat();

		// 调用方法
		dog.makeSound(); // 输出: Woof! Woof!
		cat.makeSound(); // 输出: Meow! Meow!

	Python 中多态的例子:	
		class Animal:
			def make_sound(self):
				pass

		class Dog(Animal):
			def make_sound(self):
				print("Woof! Woof!")

		class Cat(Animal):
			def make_sound(self):
				print("Meow! Meow!")

		# 实例化子类
		dog = Dog()
		cat = Cat()

		# 调用方法
		dog.make_sound()  # 输出: Woof! Woof!
		cat.make_sound()  # 输出: Meow! Meow!

*/


// 🚀 多态的例子
// 🌟 【多态参数】: 👇 s 就是一个多态参数 ————————————————————————————————————————————
type sayHi interface {
	sayHey()
}

type A struct {
	name string
}

type B struct {
	name string
}

func (a A) sayHey() {
	fmt.Printf("我是 %v \n", a.name)
}

func (b B) sayHey() {
	fmt.Printf("我是 %v \n", b.name)
}


func Greet(s sayHi) { //👈多态参数, 根据不同的【实例】来调用不同的方法
	s.sayHey()
}






func main() {
	a := A{name: "Kim"}
	b := A{name: "Zeno"}

	Greet(a)
	Greet(b)




	fmt.Println("——————————————————————————————————————————————————————————————")



	// 🌟 【多态数组】: 数组内的每个类型都是接口, 可以存放 A 的结构体, 也可以存放 B 的结构体
	var arr [2]sayHi // 🚀 都继承了 sayHi()  这个接口
	arr[0] = A{name: "Kim"}
	arr[1] = B{name: "Zeno"}
	fmt.Println(arr)

}


