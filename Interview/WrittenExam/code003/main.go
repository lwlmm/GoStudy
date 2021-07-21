package main

import(
    "fmt"
    "math"
)

func main() {
    var size1, size2 int
    var n int
	var t int
	mod := 1000000007
    fmt.Scanf("%d\n", &n)
	for i := n - 1; i >= 0; i-- {
		fmt.Scanf("%c", &t)
		t -= 96
		if t > 2 {
			t--
		}
		size1 += t * int(math.Pow(25, float64(i)))
		size1 = size1 % mod
		if size1 < 0 {
			size1 += mod
		}
	}
	fmt.Scanf("\n")
	for i := n - 1; i >= 0; i-- {
		fmt.Scanf("%c", &t)
		t -= 96
		if t > 2 {
			t--
		}
		size2 += t * int(math.Pow(25, float64(i)))
		size2 = size2 % mod
		if size2 < 0 {
			size2 += mod
		}
	}
	if size1 >= size2 {
		fmt.Println(0)
		return
	}
	ret := (size2 - size1 + 1) % mod
	if ret < 0 {
		ret += mod
	}
	fmt.Println(ret)
}