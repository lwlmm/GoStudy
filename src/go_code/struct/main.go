package main

import "fmt"

//定义结构体
type Student struct {
	name 	string
	age		int
}

//类似其他语言中的class
func (stu *Student) hello(person string) {
	fmt.Println("Hello ",person, ", I'm ", stu.name)
	return
}

func main() {

	//两种结构体实例化的方法
	stu1 := &Student{
		name: "Matt",
		age : 25,
	}
	stu1.hello("Elsa")

	//使用new
	stu2 := new(Student)
	fmt.Println("Name: ", stu2.name, ", age: ", stu2.age)
}