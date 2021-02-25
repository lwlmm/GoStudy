/*练习 7.9

给定一个slices []int 和一个 int 类型的因子factor，
扩展 s 使其长度为 len(s) * factor。*/
package main

import "fmt"

func main() {
	s := make([]int, 2, 10)
	fmt.Println("len(s) =", len(s))
	s = realloc(3, s)
	fmt.Println("len(s) =", len(s))
}

func realloc(factor int, s []int) []int {
	s = s[0: len(s) * factor]
	return s
}