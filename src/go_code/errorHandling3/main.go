package main

import(
	"fmt"
)

func get(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {

	//数组越界会触发go的panic机制
	fmt.Println(get(5))
	fmt.Println("finished")
}