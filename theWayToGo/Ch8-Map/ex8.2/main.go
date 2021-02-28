//构造一个将英文饮料名映射为法语（或者任意你的母语）的集合；
//先打印所有的饮料，然后打印原名和翻译后的名字。
//接下来按照英文名排序后再打印出来。
package main

import (
	"fmt"
	"sort"
)

func main() {
	beverage := map[string]string {
		"coke":"可乐",
		"sprite":"雪碧",
		"fenta":"芬达",
		"mountDew":"激浪",
		"doctorPepper":"屎",
	}
	fmt.Println("Unsorted list： ")
	fmt.Println(beverage)

	var list []string
	for key := range beverage {
		list = append(list, key)
	}

	sort.Strings(list)
	fmt.Println("Sorted list： ")
	for _, s := range list {
		fmt.Println(s,":", beverage[s])
	}
}