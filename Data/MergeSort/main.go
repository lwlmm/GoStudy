package main

import "fmt"

func main() {
	test := []int {5,1,1,2,0,0,108,435,24,33,551,243,93,15,4,7}
	fmt.Println(test)
	test = mergeSort(test)
	fmt.Println(test)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums)/2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) (result []int) {
	var l, r int
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}