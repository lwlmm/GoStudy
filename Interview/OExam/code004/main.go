package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	ret := make([]int, t)
	for i := 0; i < t; i++ {
		var n, percentage int
		fmt.Scanf("%d %d\n", &n, &percentage)
		a := make([]int, n)
		b := make([]int, n)
		min := math.MaxInt64 
		for j := 0; j < n; j++ {
			fmt.Scanf("%d %d\n", &a[j], &b[j])
			if min > a[j] {
				min = a[j]
			}
			if min > b[j] {
				min = b[j]
			}
		}
		var rate float64
		rate = float64(min) * float64(percentage) / 100.0
		for j := 0; j < n; j++ {
			if float64(a[j]) >= rate || float64(b[j]) >= rate {
				ret[i]++
			}
		}
	}
	for i := 0; i < t; i++ {
		fmt.Printf("Case #%d: %d\n", i+1, ret[i])
	}
}