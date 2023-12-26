package main
import (
	"fmt"
)


/* 
	Python 中的封装:
		class Person:
		def __init__(self, name; str, age: int):
			self.name = name       # 公有属性
			self.__age = age       # 私有属性(通过定义类并使用__（双下划线）前缀来创建私有成员)

		def display(self) -> None:
			# 公有方法
			print(f"Name: {self.name}, Age: {self.__age}")

		def set_age(self, age) -> None:
			# 公有方法，用于修改私有属性
			if 0 < age < 120:
				self.__age = age

		def get_age(self) -> None:
			# 公有方法，用于访问私有属性
			return self.__age

	TS 中的封装:
		class Person {
			public name: string;   // 公有属性，默认也是公有，可以省略 public
			private age: number;   // 私有属性，只能在类内部访问

			constructor(name: string, age: number) {
				this.name = name;
				this.age = age;
			}

			public display(): void {
				// 公有方法
				console.log(`Name: ${this.name}, Age: ${this.age}`);
			}

			public setAge(age: number): void {
				// 公有方法，用于修改私有属性
				if (age > 0 && age < 120) {
					this.age = age;
				}
			}

			public getAge(): number {
				// 公有方法，用于访问私有属性
				return this.age;
			}
		}

*/


func main() {
	
}


// 首字母小写, 只能在本包中使用!
type person struct {
	Name string
	Weight float32
	age int // 其他包不能访问, 可以通过工厂函数来外露这个属性
}


// 工厂函数
func NewPerson(name string) *person { // 返回指针, 避免拷贝
	return &person { // 🌟 返回局部变量的地址
		Name: name,
		Weight: 1.2, // 🔥 默认值 => 类似 js 的构造函数
	}
}


// 定义 Set 方法, 用于修改【私有属性】
func(p *person) SetAge(age int) {
	if age > 0 && age < 150 { // 年龄的合法的范围
		p.age = age
	} else {
		fmt.Println("年龄范围不正确!")
	}
}


// 定义 Get 方法, 用于访问【私有属性】
func(p *person) GetAge() int {
	return p.age
}