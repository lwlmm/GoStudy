package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    grid := make([][]int, n)
    sum := 0
    for i := 0; i < n; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scanf("%d", &grid[i][j])
            sum += grid[i][j]
        }
    }
    if sum == 0 || sum == n * n{
        fmt.Println(0)
        return
    }
    dis := make([][]int, n)
    for i := 0; i < n; i++ {
        dis[i] = make([]int, n)
    }
    ret := math.MaxInt32
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                continue
            } else {
                dist := math.MaxInt32
                det(i, j, n, grid, 0, &dist)
                dis[i][j] = dist
                if dist < ret {
                    ret = dist
                }
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if dis[i][j] == ret {
                fmt.Printf("%d %d ", i, j)
            }
        }
    }
}

func det(i, j, n int, grid [][]int,path int, ret *int) {
    if path > 2 * n {
        return
    }
    if grid[i][j] == 1 {
        if path < *ret {
            *ret = path
        }
        return
    }
    if i + 1 < n {
        det(i+1, j, n, grid, path+1, ret)
    }
    if i - 1 >= 0 {
        det(i-1, j, n, grid, path+1, ret)
    }
    if j + 1 < n {
        det(i, j+1, n, grid, path+1, ret)
    }
    if j - 1 >= 0 {
        det(i, j-1, n, grid, path+1, ret)
    }
}