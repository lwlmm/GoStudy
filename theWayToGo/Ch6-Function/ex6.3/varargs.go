//写一个函数，该函数接受一个变长参数并对每个元素进行换行打印。

package main

import "fmt"

func main() {
	printVarString("Hello", "World")
}

func printVarString(s1 ...string) {
	for _, s := range s1 {
		fmt.Println(s)
	}
	return
}