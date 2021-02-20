package main

import "fmt"

func main(){
	//仅声明 key-value
	m1 := make(map[string]int)

	//声明的同时初始化
	m2 := map[string]string{
		"Sam": "Male",
		"Alice":"Female",
	}

	//赋值/修改
	m1["Tom"] = 18

	fmt.Println(m1)
	fmt.Println(m2)
}