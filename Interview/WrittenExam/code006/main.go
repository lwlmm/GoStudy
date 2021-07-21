/*
输入描述
单组输入。

第1行输入两个正整数分别表示n和m，n和m均不超过100，两个数字之间用空格隔开。

接下来n行是一个n*m的矩阵，每一行m个正整数，表示某一叠板砖的数量，两个正整数之间用空格隔开。

输出描述
输出一个整数，即留下的水坑数量（存在一个水坑也没有的情况）。


样例输入
4 4
2 3 5 1
4 1 2 3
1 5 4 2
1 2 2 2
*/

package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	grid := make([][]int, n)
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, m)
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &grid[i][j])
		}
	}
	for i := 1; i < n - 1; i++ {
		for j := 1; j < m - 1; j++ {
			if isReal(i, j, grid) {
				ret[i][j]++
			}
		}
	} 
	cnt := 0
	for i := 1; i < n - 1; i++ {
		for j := 1; j < m - 1; j++ {
			if ret[i][j] > 0 {
				cnt++
				judge(i, j, ret)
			} else {
				continue
			}
		}
	}
	fmt.Println(cnt)
}

func judge(i, j int, grid [][]int) {
	if grid[i][j] == 0 {
		return
	}
	grid[i][j] = 0
	if i > 0 {
		judge(i-1, j, grid)
	}
	if i < len(grid) - 1 {
		judge(i+1, j, grid)
	}
	if j > 0 {
		judge(i, j-1, grid)
	}
	if j < len(grid[0])-1 {
		judge(i, j+1, grid)
	}
} 

func isReal(i, j int, grid [][]int) bool {
	ret := true
	dfs(i, j, grid, &ret, grid[i][j])
	return ret
}

func dfs(i, j int, grid [][]int, ret *bool, val int) {
	if i < 0 || i > len(grid) - 1 || j < 0 || j > len(grid[0])-1 {
		*ret = false
		return
	}
	if grid[i][j] > val {
		return
	}
	if i + 1 < len(grid) {
		dfs(i+1, j, grid, ret, grid[i+1][j])
	}
	if i  > 0 {
		dfs(i-1, j, grid, ret, grid[i-1][j])
	}
	if j > 0 {
		dfs(i, j-1, grid, ret, grid[i][j-1])
	}
	if j < len(grid[0])-1 {
		dfs(i, j+1, grid, ret, grid[i][j+1])
	}
}

func isPool(i, j int, grid [][]int) bool {
	if i <= 0 || i >= len(grid) - 1 || j <= 0 || j >= len(grid[0]) - 1 {
		return false
	}
	if  (grid[i+1][j] > grid[i][j] && grid[i-1][j] > grid[i][j] && grid[i][j+1] > grid[i][j] &&grid[i][j-1] > grid[i][j]) ||
		(isPool(i+1, j, grid) && grid[i-1][j] > grid[i][j] && grid[i][j+1] > grid[i][j] &&grid[i][j-1] > grid[i][j]) ||
		(grid[i+1][j] > grid[i][j] && isPool(i-1, j, grid) && grid[i][j+1] > grid[i][j] &&grid[i][j-1] > grid[i][j]) ||
		(grid[i+1][j] > grid[i][j] && grid[i-1][j] > grid[i][j] && isPool(i, j+1, grid) &&grid[i][j-1] > grid[i][j]) ||
		(grid[i+1][j] > grid[i][j] && grid[i-1][j] > grid[i][j] && grid[i][j+1] > grid[i][j] && isPool(i, j-1, grid)) {
			return true
	}
	return false
}