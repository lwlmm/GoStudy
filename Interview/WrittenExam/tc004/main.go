package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	ret := []int {}
	for i := 0; i < t; i++ {
		var n, charge int
		fmt.Scanf("%d %d\n", &n, &charge)
		remain := make([]int, n)
		fall := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%d %d\n", &remain[j], &fall[j])
		}
		ret = append(ret, cal(remain, fall, charge))
	}
}

func cal(r, f []int, c int) float64 {
	var ret float64
	ret = -1
	n := len(r)
	fallspeed := 0
	for i := 0; i < n; i++ {
		fallspeed += f[i]
	}
	if fallspeed <= c {
		return ret
	}
	time := make([]float64, n)
	for {
		min := float64(r[0])/float64(f[0])
		minId := 0
		for i := 0; i < n; i++ {
			time[i] = float64(r[i])/float64(f[i])
			if time[i] < min {
				minId = i
			}			
		}
		nTime := make([]float64, n)
		copy(nTime, time)
		sort.Float64(nTime)
		
	}
}