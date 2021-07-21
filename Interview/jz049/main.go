package main

import(
	"fmt"
)

func main() {
	n := 10
	ret := []int {1}
    var i1, i2, i3 int
    for {
        if len(ret) >= n {
            break
        }
        p1 := ret[i1] * 2
        p2 := ret[i2] * 3
        p3 := ret[i3] * 5
        if p1 < p2 && p1 < p3 {
            i1++
            ret = append(ret, p1)
        } else if p2 < p1 && p2 < p3 {
            i2++
            ret = append(ret, p2)
        } else if p3 < p1 && p3 < p2 {
            i3++
            ret = append(ret, p3)
        } else if
    }
    fmt.Println(ret)
}