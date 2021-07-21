package main

import (
    "fmt"
)

func main() {
    var t int
    fmt.Scanf("%d\n", &t)
    ret := []string {}
    for i := 0; i < t; i++ {
        var n int
        fmt.Scanf("%d\n", &n)
        tree := make([][]int, n)
        for j := 0; j < n; j++ {
            tree[j] = make([]int, 3)
            fmt.Scanf("%d", &tree[j][0])
            cnt := 0
            for k := 0; k < j; k++ {
                if tree[k][0] < tree[j][0] && tree[k][2] == 0 {
                    tree[k][2] = j
                    cnt++
                }
                if tree[k][0] > tree[j][0] && tree[k][1] == 0 {
                    tree[k][1] = j
                    cnt++
                }
                if cnt > 1 {
                    ret = append(ret, "NO")
                    goto end
                }
            }
        }
		ret = append(ret, "YES")
		end:
    }
	for _, s := range ret {
		fmt.Println(s)
	}
}