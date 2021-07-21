package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0,108,435,24,33,551,243,93,15,4,7}
	fmt.Println(test)
	test = popSort(test)
	fmt.Println(test)
}

func popSort(arr []int) []int {
	for i := 0; i < len(arr) - 1; i++ {
		for j := 0; j < len(arr) - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
	return arr
}