package main
import ("fmt")


/* 
	Ts 中的接口继承:
		// 定义一个接口
		interface Person {
			name: string;
			age: number;
		}

		// 定义另一个接口，继承自 Person 接口
		interface Student extends Person {
			school: string;
		}

		// 实现 Student 接口
		let student: Student = {
			name: "张三",
			age: 20,
			school: "XX大学"
		};
*/


type AInterface interface {
	aFn()
}

type BInterface interface {
	bFn()
}


type CInterface interface { // 🔥 如果实现 C 接口, 则必须实现 A 和 B 接口的所有方法
	AInterface // 继承 A 接口
	BInterface // 继承 B 接口
	cFn()
}


type Student struct {
	Age int
}


func (s Student) aFn() {
	fmt.Println("aFn")
}

func (s Student) bFn() {
	fmt.Println("bFn")
}

func (s Student) cFn() {
	fmt.Println("cFn")
}



func main() {
	studentA := Student{Age: 20}
	studentA.aFn()
}