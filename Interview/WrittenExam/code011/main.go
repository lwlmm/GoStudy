package main
import (
	"fmt"
	"math"
)

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
	ret := math.MaxInt32
	var dfs func(g [][]int, cur, dis, x, y int)
	dfs = func(g [][]int, cur, dis, x, y int) {
		if cur > k {
			if dis < ret {
				ret = dis
			}
			return
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if g[i][j] == cur {
					d := distance(i, j, x, y)
					dfs(g, cur+1, dis+d, i, j)
				}
			}
		}
	}
	dfs(grid, 1, 0, -1, -1)
    fmt.Println(ret)
}

func distance(i, j, x, y int) int {
	var a, b int
	if x == -1 && y == -1 {
		return 0
	} 
	if i > x {
		a = i - x
	} else {
		a = x - i
	}
	if j > y {
		b = j - y
	} else {
		b = y - j
	}
	return a + b
}