package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    list := make([]int, n)
    mp := make(map[int]int)
    for i := 0; i < n; i++ {
        fmt.Scanf("%d", &list[i])
        if list[i] != i + 1 {
            mp[i]++
        }
    }
    var q, sw int
    fmt.Scanf("%d", &q)
    for i := 0; i < 2; i++ {
        fmt.Scanf("%d", &sw)
        if _, ok := mp[sw-1]; ok {
            delete(mp, sw-1)
        }
    }
    if len(mp) == 0 {
        fmt.Println("YES")
        return
    }
    ret := len(mp)
    if ret % 2 == 0 {
        ret = ret / 2
    } else {
        ret = ret / 2 + 1
    }
    fmt.Println("NO")
    fmt.Println(ret)
}