package main

import(
	"fmt"
	"sort"
)

func main() {
	test := []int {1, 2, 4, 8, 16}
	fmt.Println(largestDivisibleSubset(test))
}

func largestDivisibleSubset(nums []int) []int {
    sort.Ints(nums)
    dp := [][]int {}
    for i := 0; i < len(nums); i++ {
        tmp := [][]int {}
        tmp = append(tmp, []int {nums[i]})
        for j := 0; j < len(dp); j++ {
            if nums[i] % dp[j][len(dp[j])-1] == 0 {
                t := append(dp[j], nums[i])
                tmp = append(tmp, t)
            }
        }
        dp = append(dp, tmp...)
		fmt.Println(dp)
    }
    max := 0
    ret := 0
    for i, d := range dp {
        if len(d) > max {
            max = len(d)
            ret = i
        }
    }
    return dp[ret]
}