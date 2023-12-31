package main
import ("fmt")


/*
	What?
		抽象类, 让子类去继承公共的【属性】或【方法】

	How?
		在结构体内定义【匿名结构体】


	TS 中的继承:
		// 定义一个抽象类
		abstract class Animal {
			// 定义一个抽象方法，没有具体实现，必须在派生类中实现
			abstract makeSound(): void;

			// 定义一个普通方法，有具体的实现
			move(): void {
				console.log("roaming the earth...");
			}
		}

		// 定义一个派生类，继承自Animal
		class Dog extends Animal {
			// 实现父类中的抽象方法
			makeSound() {
				console.log("Woof! Woof!");
			}
		}

	Python 中的继承:
		from abc import ABC, abstractmethod

		class Animal(ABC):
			@abstractmethod
			def move(self):
				pass

		class Human(Animal):
			def move(self):
				print("我可以走和跑")

		class Fish(Animal):
			def move(self):
				print("我可以游泳")
*/

func main() {
	cat := &Cat{
		Animal: Animal{
			Name: "小猫",
			Weight: 1.2,
			Age: 2,
		},
		Age: 3,
	}

	fmt.Println(*cat)
	fmt.Println(cat.Animal.Age) // 公共属性 2
	fmt.Println(cat.Age) // 子类自己的属性 3
	fmt.Println(cat.Animal.Name) // 公共属性 "小猫"
	cat.Age = 4 // 改变子类自己的属性
	fmt.Println(cat.Age) // 子类自己的属性 4
	cat.Animal.Bark() // 公共方法
	cat.Animal.ShowInfo() // 公共方法 => 🔥 也可以 cat.ShowInfo() 直接访问 => 如果子类没有此方法, 则会调用父类的方法 (也要看子类有没有重写此方法)
	cat.Scratch() // 子类自己的方法
}

// 父类
type Animal struct {
	Name string
	Weight float32
	Age int
}


func (ani *Animal) Bark() { // 公共方法
	fmt.Println("喊叫-公共方法")
}

func (ani *Animal) ShowInfo() { // 公共方法
	age := ani.Age
	weight := ani.Weight
	fmt.Printf("展示信息公共方法-年龄: %d, 体重: %f\n", age, weight)
}


// 子类
type Cat struct {
	Animal // 匿名结构体, 继承 Animal的【属性】和【方法】
	Age int
}

func (cat *Cat) Scratch() { // 子类自己的方法!!
	fmt.Println("抓人-子类自己的方法")
}