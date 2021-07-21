package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0,108,435,24,33,551,243,93,15,4,7}
	fmt.Println(test)
	test = insertSort(test)
	fmt.Println(test)
}

func insertSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j >= 1 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}