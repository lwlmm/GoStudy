package main

import "fmt"

type Person interface {
	getName()	string
	getAge()	int
}
//结构必须实现接口的所有方法才能implement
//否则会编译错误

type Student struct {
	name 	string
	age		int
}

func (stu *Student) getAge() int {
	return stu.age
}

func (stu *Student) getName() string {
	return stu.name
}

func main(){

	//接口和结构可以互相转换
	stu1 := &Student{
		name:	"Matt",
		age:	25,
	}
	var p Person = stu1		//实例转换为接口
	fmt.Println(p.getName())
	
	stu2 := p.(*Student)	//将接口转换为实例
	stu2.getAge()

	//为确保某个类型实现某个接口的所有方法，可以用nil强制转换进行检测
	var _ Person = (*Student)(nil)

	//如果定义了一个没有方法的空接口，这个接口可以表示任何类型
	m := make(map[int]interface{})
	m[0] = "String"
	m[1] = 1
	m[2] = 3.14
	m[3] = [4]string{"This ", "is ", "empty", "interface"}
	fmt.Println(m)
}


