package main

import "fmt"

func main(){

	//对数组，切片，字典使用 for range 遍历
	nums := []int{10, 20, 30, 40}
	for i, num := range nums{
		fmt.Println(i, num)
	}

	slice1 := nums[1:4]
	for i, num := range slice1{
		fmt.Println(i, num)
	}

	m2 := map[string]string{
		"Sam": "Male",
		"Amy": "Female",
	}

	for key, value := range m2{
		fmt.Println(key, value)
	}
}