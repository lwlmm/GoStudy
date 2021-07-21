package main

import(
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	seat := make([]int, 24)
	for i := 0; i < n; i++ {
		var start, end, num int
		fmt.Scanf("%d %d %d\n", &start, &end, &num)
		flag := false
		for j := start; j < end; j++ {
			if seat[j] + num > m {
				flag = true
				break
			}
		}
		if !flag {
			for j := start; j < end; j++ {
				seat[j] += num
			}
		}
	}
	for i := 0; i < 24; i++ {
		fmt.Printf("%d ", seat[i])
	}
}