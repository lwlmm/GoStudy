package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0,108,435,24,33,551,243,93,15,4,7}
	fmt.Println(test)
	test = quickSort(test)
	fmt.Println(test)
}

func quickSort(nums []int) []int {
	qsort(nums, 0, len(nums)-1)
	return nums
}

func qsort(nums []int, start, end int) {
	if start < end {
		pivot := partition(nums, start, end)
		qsort(nums, start, pivot - 1)
		qsort(nums, pivot + 1, end)
	}
}

func partition(nums []int, start, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}