package main

import(
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Scanf("%d %d", &n, &m)
		s := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &s[j])
		}
		fmt.Println(cmp(s, n, m))
	}
}

func cmp(s []int, n, m int) int {
	ret := -1
	cur := n+1
	sort.Ints(s)
	dfs(s, n, m, 0, &cur)
	if cur == n + 1 {
		return ret
	}
	return cur
}

func dfs(s []int, n, m, sum int, cur *int) {
	for i := 1; ; i++ {
		if i * m > sum {
			break
		}
		if i * m == sum {
			if n - len(s) < *cur {
				*cur = n - len(s)
			}
		}
	}
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s); i++ {
		t := sum + s[i]
		dfs(cut(s, i), n, m, t, cur)
	}
}

func cut(s []int, i int) []int {
	ret := []int {}
	for j := 0; j < len(s)-1; j++ {
		if j != i {
			ret = append(ret, s[j])
		}
	}
	return ret
}