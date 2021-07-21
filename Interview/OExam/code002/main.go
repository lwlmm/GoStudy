package main

import(
	"fmt"
	"math/rand"
	"time"
)

var SIZE int = 100
var SCALE int = 500

func main() {
	rand.Seed(time.Now().Unix())
	test := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		test[i] = rand.Intn(SCALE)
	}
	test = bucketSort(test)
	fmt.Println(test)
}

func bucketSort(arr []int) []int {
	num := len(arr)
	max := getMax(arr)
	buckets := make([][]int, num)			//开辟桶的数量与数列长度一致

	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num-1) / max		//分配桶编号：index = value * (n-1) / k
		buckets[index] = append(buckets[index], arr[i])		//入桶
	}

	tmpPosition := 0						//记录返回数列的标号

	for i := 0; i < num; i++ {
		bucketLength := len(buckets[i])
		if bucketLength > 0 {
			sortInBucket(buckets[i])
			copy(arr[tmpPosition:], buckets[i])	//将排序后的桶返回
			tmpPosition += bucketLength
		}
	}

	return arr
}

func getMax(arr []int) int {
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}

func sortInBucket(arr []int) {
	arr = quickSort(arr)
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
