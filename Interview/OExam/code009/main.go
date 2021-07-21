package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	ret := []int {}
	client := make([][]int, n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scanf("%d\n", &t)
		client[i] = make([]int, k)
		for j := 0; j < k; j++ {
			client[i][j] = t * (j+1)
		}
	}
	cnt := 0
	for {
		for j := 0; j < n; j++ {
			if client[j][0] == cnt {
				client[j] = client[j][1:]
				ret = append(ret, j+1)
				cnt--
				break
			}
		}
		cnt++
		if len(ret) == k {
			break
		}
	}
	for _, c := range ret {
		fmt.Println(c)
	}
}