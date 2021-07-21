package main

import (
	"fmt"
	"math"
)

func main() {
	house := []int {3,1,2,3}
	cost := [][]int {{1,1,1},{1,1,1},{1,1,1},{1,1,1}}
	fmt.Println(minCost(house, cost, 4, 3, 3))
}

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
    dp := make([][][]int, m)
    for i := 0; i < m; i++ {
        dp[i] = make([][]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = make([]int, target)
            for k := 0; k < target; k++ {
                dp[i][j][k] = math.MaxInt32
            }
        }
    }
    if houses[0] != 0 {
        dp[0][houses[0]-1][0] = 0
    } else {
        for i := 0; i < n; i++ {
            dp[0][i][0] = cost[0][i]
        }
    }
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            for k := 0; k < target; k++ {
                if houses[i] != 0 {
                    if houses[i] - 1 == j {
                        dp[i][j][k] = dp[i-1][j][k]
                    } else if k > 0 {
                        min := math.MaxInt32
                        for l := 0; l < n; l++ {
                            if l != j {
                                if dp[i-1][l][k-1] < min {
                                    min = dp[i-1][l][k-1]
                                }
                            }
                        }
                        dp[i][j][k] = min
                    }
                    continue
                }
                min := math.MaxInt32
                if k > 0 {
                    for l := 0; l < n; l++ {
                        if l != j {
                            if dp[i-1][l][k-1] < min {
                                min = dp[i-1][l][k-1]
                            }
                        }
                    }
                }
                if dp[i-1][j][k] < min {
                    min = dp[i-1][j][k]
                }
                if min == math.MaxInt32 {
                    dp[i][j][k] = min
                } else {
                    dp[i][j][k] = min + cost[i][j]
                }
            }
        }
    }
    ret := math.MaxInt32
    for i := 0; i < n; i++ {
        if dp[m-1][i][target-1] < ret {
            ret = dp[m-1][i][target-1]
        }
    }
	fmt.Println(dp)
    if ret == math.MaxInt32 {
        return -1
	}
    return ret
}