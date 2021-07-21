package main
import "fmt"

func main() {
    var n, k int
    fmt.Scanf("%d %d\n", &n, &k)
    list := make([]int, k+1)
    grid := make([][]int, n)
    for i := 0; i < n; i++ {
        grid[i] = make([]int, n)
        for j := 0; j < n; j++ {
            fmt.Scanf("%d", &grid[i][j])
            if grid[i][j] <= k {
                list[grid[i][j]]++
            }
        }
    }
    for i := 1; i <= k; i++ {
        if list[i] == 0 {
            fmt.Println(-1)
            return
        }
    }
    fmt.Println(k-1)
}