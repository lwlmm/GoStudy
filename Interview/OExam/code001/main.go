package main

import(
	"fmt"
	"math/rand"
	"time"
)

var SIZE int = 100
var SCALE int = 1000

func main() {
	rand.Seed(time.Now().Unix())
	test := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		test[i] = rand.Intn(SCALE)
	}
	test = quickSort(test)
	fmt.Println(test)
}

func quickSort(arr []int) []int {
	qsort(arr, 0, len(arr)-1)
	return arr
}

func qsort(arr []int, start, end int) {
	if start < end {
		pivot := partition(arr, start, end)
		qsort(arr, start, pivot-1)
		qsort(arr, pivot+1, end)
	}
}

func partition(arr []int, start, end int) int {
	p := arr[end]
	i := start
	for j := i; j < end; j++ {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return i
}
