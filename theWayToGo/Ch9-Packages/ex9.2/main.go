//通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。

package main

import (
	"fmt"
	"unsafe"
)

func main()  {
	fmt.Printf("size of int is := %d\n", unsafe.Sizeof(1))
}