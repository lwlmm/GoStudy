package main

import(
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	ret := []string {}
	for i := 0; i < t; i++ {
		var sa, sb string
		fmt.Scanf("%s\n%s\n", &sa, &sb)
		if isSame(sa, sb) {
			ret = append(ret, "YES")
		} else {
			ret = append(ret, "NO")
		}
	}
	for _, ss := range ret {
		fmt.Println(ss)
	}
}

func isSame(s1, s2 string) bool {
	a := len(s1)
	b := len(s2)
	if a != b {
		return false
	}
	if a % 2 == 1 {
		for i := 0; i < a; i++ {
			if s1[i] != s2[i] {
				return false
			}
		}
		return true
	}
	mid := a / 2
	return (isSame(s1[:mid], s2[:mid])&& isSame(s1[mid:], s2[mid:])) || (isSame(s1[:mid], s2[mid:])&& isSame(s1[mid:], s2[:mid]))
}