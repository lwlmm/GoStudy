package main

import "fmt"

func main() {
	var arr [3][5]int
	fmt.Println(len(arr))
	fmt.Println(len(arr[len(arr)-1]))
	for i := 0; i < 3/2; i++ {
		fmt.Println(i)
	}
}