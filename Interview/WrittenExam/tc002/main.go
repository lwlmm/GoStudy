package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	var out string
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &s[j])
		}
		tmp := fmt.Sprintf("%d\n", cal(s))
		out += tmp
	}
	fmt.Printf(out)
}

func cal(s []int) int {
	n := len(s)
	if n == 1 {
		return s[0]
	}
	if n == 2 {
		return s[0]+s[1]
	}
	tail := 0
	if n % 2 == 1 {
		tail = 1
	}
	group := (n - tail - 2) / 2
	sort.Ints(s)
	ret := 0
	for i := 0; i < group; i++ {
		ret += (s[1] * 2 + s[0] + s[(i+1)*2+1])
	}
	if tail == 0 {
		return ret + s[1]
	}
	return ret + s[0] + s[1] + s[len(s)-1]
}