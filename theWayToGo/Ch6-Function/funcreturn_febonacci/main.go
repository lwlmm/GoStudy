package main

import "fmt"

func fib() func(int) int {
	var x, y int
	return func(a int) int {
		if x == 0 {
			x = 1
		}
		if y == 0 {
			y = 1
		}
		t := x
		x += y
		y = t
		return x
	}
}

func fibb(no int) int {
	if(no == 1 || no == 2) {
		return 1
	}
	var f = fib()
	j := 1
	for i := 2; i < no; i++ {
		j = f(j)
	}
	return j
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fibb(i))
	}
}