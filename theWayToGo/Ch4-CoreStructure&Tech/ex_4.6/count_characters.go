//创建一个用于统计字节和字符（rune）的程序，
//并对字符串 asSASA ddd dsjkdsjs dk 进行分析，
//然后再分析 asSASA ddd dsjkdsjsこん dk，
//最后解释两者不同的原因（提示：使用 unicode/utf8 包）。

package main

import (
	"fmt"
	"unicode/utf8"
)

func countByte(s string) {
	fmt.Println("The number of bytes in string: '",
				 s, "', is: ", len(s))
}

func countRune(s string) {
	fmt.Println("The number of bytes in string: '",
				 s, "', is: ", utf8.RuneCountInString(s))
}

func main()  {
	str1 := "asSASA ddd dsjkdsjs dk"
	str2 := "asSASA ddd dsjkdsjsこん dk"

	countByte(str1)
	countRune(str1)
	
	countByte(str2)
	countRune(str2)
}