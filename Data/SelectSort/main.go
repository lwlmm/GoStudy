package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0,108,435,24,33,551,243,93,15,4,7}
	fmt.Println(test)
	test = selectSort(test)
	fmt.Println(test)
}

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		t := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[t] {
				t = j
			}
		}
		arr[i], arr[t] = arr[t], arr[i]
	}
	return arr
}