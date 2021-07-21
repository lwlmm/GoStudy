/*

输入描述
第一行A，B两个整数；

第二行A个整数a1,a2,a3,...,aA，表示小团手上第i颗糖果的甜度；

第三行B个整数b1,b2,b3,...,bB，表示小美手上第i颗糖果的甜度。

1<=A+B<=200000, -100<=ai,bi<=100。

输出描述
输出仅包含一个非负整数，表示能够获得的甜度之和的最大值，例如当糖果都为负数的时候，答案为0。

样例输入
3 4
1 -1 0
*/

package main

import(
	"fmt"
)

func main() {
	var a, b int
	fmt.Scanf("%d %d\n", &a, &b)
	su1 := make([]int, a)
	su2 := make([]int, b)
	for i := 0; i < a; i++ {
		fmt.Scanf("%d", &su1[i])
	}
	for i := 0; i < b; i++ {
		fmt.Scanf("%d", &su2[i])
	}
	fmt.Println(max(su1)+max(su2))
}

func max(s []int) int {
	sum := 0
	cur := 0
	for _, r := range s {
		cur += r
		if cur > sum {
			sum = cur
		}
	}
	return sum
}