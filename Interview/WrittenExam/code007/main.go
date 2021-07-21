package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	s := make([]int, n)
	var t byte
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &t)
		if t == 'a' {
			s[i] = 0
		} else {
			s[i] = 1
		}
	}
	fmt.Scanf("\n")
	left, right, ret := 0, 0, 0
	cnt := 0
	cntN := m
	for right < n {
		if s[right] == 0 {
			cnt++
		} else {
			cntN--
			if cntN >= 0 {
				cnt++
            } else {
                cnt++
                for s[left] != 1 {
				    left++
				    cnt--
			    }
			    left++
			    cnt--
			    cntN++
            }
		}
		if cnt > ret {
			ret = cnt
		}
		right++
	}
	left, right, cnt, cntN = 0, 0, 0, m
	for right < n {
		if s[right] == 1 {
			cnt++
		} else {
			cntN--
			if cntN >= 0 {
				cnt++
            } else {
                cnt++
                for s[left] != 0 {
				    left++
				    cnt--
			    }
			    left++
			    cnt--
			    cntN++
            }
		}
		if cnt > ret {
			ret = cnt
		}
		right++
	}
	fmt.Println(ret)
}
