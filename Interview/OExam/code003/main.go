package main

import(
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	price := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &price[i])
	}
	ret := math.MaxInt64
	var dfs func(s []int, cur int)
	dfs = func(s []int, cur int) {
		if cur >= m {
			if cur < ret {
				ret = cur
			}
			return
		}
		for i := 0; i < len(s); i++ {
			dfs(s[i+1:], cur+s[i])
		}
	}
	dfs(price, 0)
	fmt.Println(ret)
}