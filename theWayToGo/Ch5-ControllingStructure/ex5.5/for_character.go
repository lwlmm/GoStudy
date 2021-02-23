/*创建一个程序，要求能够打印类似下面的结果（尾行达 25 个字符为止）：
G
GG
GGG
GGGG
GGGGG
GGGGGG
*/
package main

import "fmt"

func main() {
	var s string
	for i := 1; i <= 25; i++ {
		s += "G"
		fmt.Println(s)
	}
}