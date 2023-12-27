package main
import "fmt"
/* 
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

func main() {

}