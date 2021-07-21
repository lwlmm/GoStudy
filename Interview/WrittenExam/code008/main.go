package main

import(
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	cnt := 1
	var dfs func(s, cur string)
	m := make(map[byte]bool)
	dfs = func(s, cur string) {
		n := len(s)
		for i := 0; i < n; i++ {
			if /*isUnique(cur, s[i])*/ _, ok := m[s[i]]; !ok {
				m[s[i]] = true
				dfs(string(s[i+1:]), cur+string(s[i:i+1]))
				cnt++
				cnt = cnt % 20210101
				delete(m, s[i])
			}
		}
	}
	var cur string
	dfs(s, cur)
	fmt.Println(cnt)
}

func cut(s string, i int) string {
	left := s[:i]
	if i == len(s) - 1 {
		return string(left)
	}
	right := s[i+1:]
	return string(left+right)
}

func isUnique(s string,c byte) bool {
	for _, r := range s {
		if byte(r) == c {
			return false
		}
	}
	return true
}