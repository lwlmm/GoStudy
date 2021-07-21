package main

import(
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	a := make([]int, n)
	b := make([]int, n)
	fruit := make(map[int]int)
	day := 0
	ret := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &b[i])
		t := i + b[i] - 1
		if _, ok := fruit[t]; ok {
			fruit[t] += a[i]
		} else {
			fruit[t] = a[i]
		}
		if t > day {
			day = t
		}
		for j := i; j <= day; j++ {
			if _, ok := fruit[j]; ok {
				if fruit[j] > 0 {
					fruit[j]--
					ret++
					break
				}
			}
		}
	}
	for i := n; i <= day; i++ {
		for j := i; j <= day; j++ {
			if _, ok := fruit[j]; ok {
				if fruit[j] > 0 {
					fruit[j]--
					ret++
					break
				}
			}
		}
	}
	fmt.Println(ret)
}