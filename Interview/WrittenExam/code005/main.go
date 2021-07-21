/*
小明最近在学习《有机化学》。在化学中，有机化合物简称为有机物，是指含碳元素的化合物，其组成元素除碳外，通常还含有氢、氧、氮、硫、磷、卤素等元素。

小明想写个小程序来计算简单有机化合物的分子量（相对分子量），假设需要计算分子量的有机物中只包含C、H、O、N四种元素，其相对原子质量分别如下：

C: 12，H: 1，O：16，N: 14。

例如某有机物的分子式为：C4H9OH，其相对分子量为12*4+1*9+16+1 = 74。

请你帮他实现这个计算相对分子量的程序。
*/

package main

import (
	"fmt"
)

func main() {
	ret := 0
	var ele, num int
	var dic byte
	fromNum := true
	for {

        _, err := fmt.Scanf("%c", &dic)
        if err != nil || dic == '\n' {
            break
        }
		if fromNum {
			if isEle(dic) {
				ret += ele * num
				num = 0
				ele = getNum(int(dic))
				fromNum = false
			} else {
				num *= 10
				num += int(dic - '0')
			}
		} else {
			if isEle(dic) {
				ret += ele
				ele = getNum(int(dic))
			} else {
				num += int(dic - '0')
				fromNum = true
			}
		}
    }
	if fromNum {
		ret += ele * num
	} else {
		ret += ele
	}
	fmt.Println(ret)
}

func isEle(dic byte) bool {
	if dic == 'C' || dic == 'H' || dic == 'O' || dic == 'N' {
		return true
	}
	return false
}

func getNum(ele int) int {
	switch  ele {
		case 67:
			return 12
		case 72:
			return 1
		case 79:
			return 16
		case 78:
			return 14
	}
	return 0
}
