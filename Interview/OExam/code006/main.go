package main

import(
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	ret := []string {}
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &arr[j])
		}
		if increase3(arr) {
			ret = append(ret, "Yes")
		} else {
			ret = append(ret, "No")
		}
	}
	for _, s := range ret {
		fmt.Println(s)
	}
}

func increase3(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	max1, min2 := math.MinInt64, math.MaxInt64
	for _, v := range arr {
		if max1 <= v {
			max1 = v
		} else if min2 >= v {
			min2 = v
		} else {
			return true
		}
	}
	return false
}