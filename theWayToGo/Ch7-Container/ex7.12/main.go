//编写一个函数，要求其接受两个参数，
//原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。
package main

import "fmt"

func main() {
	s := "HelloWorld"
	fmt.Println(cut(s, 5))
}

func cut(s string, i int) (string, string) {
	s1 := s[:i]
	s2 := s[i:]
	return s1, s2
}