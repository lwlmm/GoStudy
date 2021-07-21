package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    side := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &side[i])
    }
    sum := make([]int, n)
    largest := 0
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            for k := j+1; k < n; k++ {
                if isTri(side[i], side[j], side[k]) {
                    sum[i]++
                    sum[j]++
                    sum[k]++
                    if sum[i] > largest {
                        largest = sum[i]
                    }
                    if sum[j] > largest {
                        largest = sum[j]
                    }
                    if sum[k] > largest {
                        largest = sum[k]
                    }
                }
            }
        }
    }
    out := []int {}
    for i := 0; i < n; i++ {
        if sum[i] == largest {
            if notIn(side[i], out) {
                out = append(out, side[i])
            }
        }
    }
    for i := 0; i < len(out); i++ {
        fmt.Printf("%d ", out[i])
    }
    fmt.Printf("\n")
}

func notIn(n int, out []int) bool {
    for _, r := range out {
        if r == n {
            return false
        }
    }
    return true
}

func isTri(i, j, k int) bool {
    if i+j <= k || j+k <= i || i+k <= j {
        return false
    }
    return true
}