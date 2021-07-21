/*
小美发明了一个函数：f(x)，表示将x的所有正整数因子提取出来之后从小到大排列，再将数字拼接之后得到的数字串。例如：10的所有因子为1,2,5,10，那么将这些因子从小到大排序之后再拼接得到的数字串为12510，即f(10)=12510。

小美十分讨厌数字k，如果f(x)中含有某个子串对应的数字等于k，那么她的不高兴度就会增加1。例如：f(8)=1248，那么其所有的子串为：1,2,4,8,12,24,48,124,248,1248，只要k等于其中任意一个值，那么小美的不高兴度就会增加1。对于每一个数，其不高兴度至多增加1。

现在，有一个长度为n的正整数序列，定义其不高兴度为序列中所有数的不高兴度的总和。小美想要知道自己对于这个序列的总不高兴度为多少。



输入描述
第一行两个正整数n,k；

第二行n个正整数ai，表示该序列。

1≤n≤105，1≤k，ai≤2x104
*/
package main

import(
	"fmt"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	kk := strconv.Itoa(k)
	cnt := 0
	var test int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &test)
		s := combine(test)
		if isUg(s, kk) {
			cnt++
		}
	}
	fmt.Println(cnt)
}

func isUg(s, t string) bool {
	for i := 0; i <= len(s) - len(t); i++ {
		for j := 0; j < len(t); j++ {
			if t[j] != s[i+j] {
				goto lab
			}
		}
		return true
		lab:
	}
	return false
}

func combine(t int) string {
	var s string
	for i := 1; i <= (t/2)+1; i++ {
		if t % i == 0 {
			tmp := strconv.Itoa(i)
			s += tmp
		}
	}
	tmp := strconv.Itoa(t)
	s += tmp
	return s
}