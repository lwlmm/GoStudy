package main

import (
    "fmt"
    "math"
)

func main() {
    var x, y, t int
    fmt.Scanf("%d\n", &x)
    fmt.Scanf("%d\n", &y)
    price := []int {}
    for {
        _, err := fmt.Scanf("%d", &t)
        if err != nil {
            break
        }
        price = append(price, t)
    }
    ret := math.MaxInt32
    dfs(price, 0, x, -1, &ret)
    fmt.Println(ret-y)
}

func dfs(price []int, sum, x, i int, ret *int) {
	if sum > x {
		if sum < *ret {
		    *ret = sum
		}
		return
	}
	for j := i + 1; j < len(price); j++ {
		dfs(price, sum+price[j], x, j, ret)
	}
}