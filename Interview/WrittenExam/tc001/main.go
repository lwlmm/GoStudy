package main

import(
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%d\n", &n)
    s := make([]int, n)
    var c byte
    for i := 0; i < n; i++ {
        fmt.Scanf("%c", &c)
        s[i] = int(c-'0')
    }
    fmt.Println(sum(s))
}

func sum(s []int) int {
	if len(s) < 2 {
		return len(s)
	}
	for i := 0; i < len(s) - 1; i++ {
		if s[i] + s[i+1] == 10 {
			if i == 0 {
				t := []int {}
				t = append(t, s[i+2:]...)
				return sum(t)
			} else {
				t := []int {}
				t = append(t, s[i-1])
				t = append(t, s[i+2:]...)
				return i - 1 + sum(t)
			}
		}
	}
	return len(s)
}

func dfs(s *[]int) {
    if len(*s) < 2 {
        return
    }
    for i := 0; i < len(*s)-1; i++ {
        if (*s)[i] + (*s)[i+1] == 10 {
            if i == len(*s) - 2 {
                *s = (*s)[:i]
                dfs(s)
            } else {
                left := (*s)[:i]
                *s = append(left, (*s)[i+2:]...)
                dfs(s)
            }
        }
    }
}