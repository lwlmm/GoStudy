//创建一个 map 来保存每周 7 天的名字，
//将它们打印出来并且测试是否存在 Tuesday 和 Hollyday。

package main

import "fmt"

func main() {
	Week := make(map[string]int)
	Week["Monday"] = 1
	Week["Tuesday"] = 2
	Week["Wednesday"] = 3
	Week["Thursday"] = 4
	Week["Friday"] = 5
	Week["Saturday"] = 6
	Week["Sunday"] = 7

	if _, ok := Week["Hollyday"]; ok {
		fmt.Println("Hollyday is in a Week!")
	} else {
		fmt.Println("Hollyday is not in a week")
	}
	fmt.Println(Week)
}