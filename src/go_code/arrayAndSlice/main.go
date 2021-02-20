package main

import "fmt"

func main(){
	var arr = [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++{
		arr[i] += 100
	}
	fmt.Println(arr)

	slice1 := make([]float32, 0)	//slice of 0 length
	slice2 := make([]float32, 3, 5)	//[0 0 0] 长度为3容量为5的切片
	fmt.Println(slice1)
	fmt.Println(slice2, len(slice2), cap(slice2))

	//切片使用
	slice2 = append(slice2, 1, 2, 3)	//切片容量可以根据需要自动扩展
	fmt.Println(len(slice2), cap(slice2))

	//子切片
	sub1 := slice2[3:]	// [1 2 3] 第3个以后
	sub2 := slice2[:3] 	// [0 0 0]
	sub3 := slice2[1:4]	// [0 0 1]

	// 合并切片
	combined := append(sub1, sub2...)

	fmt.Println(sub1)
	fmt.Println(sub2)
	fmt.Println(sub3)
	fmt.Println(combined)
}