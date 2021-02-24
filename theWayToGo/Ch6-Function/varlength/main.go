//如果函数的最后一个参数是采用 ...type 的形式，
//那么这个函数就可以处理一个变长的参数，
//这个长度可以为 0，这样的函数称为变参函数。

package main

import "fmt"

func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func min(s ...int) int {
	if len(s)==0 {
		return 0
	}
	min := s[0]
	for _, v := range s {	//使用for range 来遍历slice
		if v < min {
			min = v
		}
	}
	return min
}