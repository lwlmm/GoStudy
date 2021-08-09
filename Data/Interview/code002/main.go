package main

import (
    "fmt"
    "sort"
)

func main() {
    var n, m int
    fmt.Scanf("%d %d\n", &n, &m)
    nums := make([]int, n)
    for i := range nums {
        fmt.Scanf("%d", &nums[i])
    }
    sort.Ints(nums)
    ret := 0
    flags := make([]bool, n)
    var dfs func(cur []int, flag []bool)
    dfs = func(cur []int, flag []bool) {
        if len(cur) == n {
            ret++
            return
        }
        for i, num := range nums {
            if flag[i] {
                continue
            }
            if i > 0 && nums[i] == nums[i-1] && !flag[i-1] {
                continue
            }
            if cur[len(cur)-1] <= num + m {
                flag[i] = true
                dfs(append(cur, num), flag)
                flag[i] = false
            }
        }
    }
    dfs([]int {}, flags)
    fmt.Println(ret)
}