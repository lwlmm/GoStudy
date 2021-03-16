package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0}
	fmt.Println(test)
	HeapSort(test)
	fmt.Println(test)
}

/*

*/
func HeapSort(arr []int) {
	buildMaxHeap(arr)
	for i := len(arr)-1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, 0, i)
	}
}

func buildMaxHeap(arr []int) {
	for i := len(arr) / 2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr))
	}
}

func maxHeapify(arr []int, i, length int) {
	l := 2 * i + 1
	r := 2 * i + 2
	max := i
	if l < length && arr[l] > arr[max] {
		max = l
	}
	if r < length && arr[r] > arr[max] {
		max = r
	}
	if max == i {
		return
	} else {
		arr[i], arr[max] = arr[max], arr[i]
		maxHeapify(arr, max, length)
	}
}

