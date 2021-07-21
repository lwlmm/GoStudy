package main

import "fmt"

func main() {
	s := "a"
	t := "a"
	dp := make([][]int, len(t)+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, len(s)+1)
    }
    for j := 1; j < len(dp[0]); j++ {
        dp[0][j] = 1
    }
	fmt.Println(dp)
}