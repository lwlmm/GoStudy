package main

import(
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	res := []int {}
	for i := 0; i < t; i++ {
		var n int
		fmt.Scanf("%d\n", &n)
		time := make([][]int, n)
		for j := 0; j < n; j++ {
			time[j] = make([]int, 2)
			fmt.Scanf("%d", &time[j][0])
		}
		for j := 0; j < n; j++ {
			fmt.Scanf("%d", &time[j][1])
		}
		sort.Slice(time, func(i, j int) bool {
			if time[i][0] < time[j][0] {
				return true
			} else if time[i][0] == time[j][0] {
				return time[i][1] < time[j][1]
			} else {
				return false
			}
		})
		ret := 0
		for j := time[n-1][0]; j >= 1; j-- {
			curTime := -1
			for k := len(time)-1; k >= 0; k-- {
				if time[k][0] >= j {
					curTime = k
					break
				}
			}
			time = cut(time, curTime)
		}
		for _, t := range time {
			ret += t[1]
		}
		res = append(res, ret)
	}
	for _, cnt := range res {
		fmt.Println(cnt)
	}
}

func cut(time [][]int, cur int) [][]int {
	ret := [][]int {}
	for i := 0; i < len(time); i++ {
		if i != cur {
			ret = append(ret, time[i])
		}
	}
	return ret
}