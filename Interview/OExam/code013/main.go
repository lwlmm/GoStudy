package main

import(
	"fmt"
)

func main() {
	var m, n int
	fmt.Scanf("%d\n%d\n", &m, &n)
	pool := make([][]int, m)
	mod := make([]int, m)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		pool[a] = append(pool[a], b)
		mod[a] = 1
		mod[b] = 1
	}
	ret := 10
	var isCircle func(i int) bool 
	isCircle = func(i int) bool {
		mp := make(map[int]bool)
		stack := []int {}
		for j := 0; j < len(pool[i]); j++ {
			stack = append(stack, pool[i][j])
		}
		mp[i] = true
		for len(stack) > 0 {
			tmp := []int {}
			for _, s := range stack {
				if _, ok := mp[s]; ok {
					return true
				} else {
					mp[s] = true
				}
				for j := 0; j < len(pool[s]); j++ {
					tmp = append(tmp, pool[s][j])
				}
			}
			stack = append([]int(nil), tmp...)
		}
		return false
	}

	for _, v := range mod {
		if v != 1 {
			ret--
		}
	}
    for i, _ := range mod {
        if isCircle(i) {
            ret -= 2
            break
        }
    }
	if ret >= 0 {
		fmt.Println(ret)
	} else {
		fmt.Println(0)
	}
}