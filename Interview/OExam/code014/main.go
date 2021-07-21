package main

import(
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Scanf("%d\n", &k)
	cost := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d\n", &cost[i])
	}
	var n, e int
	fmt.Scanf("%d %d\n", &n, &e)
	pool := make([][]int, n)
	for i := 0; i < e; i++ {
		var f, t int
		fmt.Scanf("%d %d\n", &f, &t)
		pool[f] = append(pool[f], t)
		pool[t] = append(pool[t], f)
	}
	ret := math.MaxInt64
	for i := 0; i < k; i++ {
		testHot(pool, cost[i], i+2, n, &ret )
	}
	fmt.Println(ret)
}

func testHot(pool [][]int, cost, length, n int, ret *int) {
	hot := make([]int, n)
	for i := 0; i < 2; i++ {
		dfs(pool, hot, cost, length, i+1, 0, n, ret)
	}
}

func dfs(pool [][]int, hot []int, cost, length, time, round, n int, ret *int) {
	cur := cost * time
	for curin := 0; curin < len(pool); curin++ {
		hottie := make([]int, n)
		copy(hottie, hot)
		stack := []int {curin}
		for j := 0; j < length; j++ {
			tmp := []int {}
			for _, ele := range stack {
				hottie[ele] = 1
				for _, related := range pool[ele] {
					tmp = append(tmp, related)
				}
			}
			stack = append([]int(nil), tmp...)
		}
		hotted := true
		for _, ishot := range hottie {
			if ishot != 1 {
				hotted = false
				break
			}
		}
		if hotted {
			if cur < *ret {
				*ret = cur
			}
			break
		}
		pool[curin] = []int {}
		round++
		if round < time {
			dfs(pool, hottie, cost, length, time, round, n, ret)
		}
	}
}