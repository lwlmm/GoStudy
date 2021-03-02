//使用 container/list 包实现一个双向链表，
//将 101、102 和 103 放入其中并打印出来。

package main

import (
	"fmt"
	"container/list"
)

func main() {
	l := &list.List{}
	i := 101
	for e := l.Front(); e != nil; e = e.Next() {
		e.Value = i
		i++
		if i == 104 {
			break
		}
	}
	fmt.Println(l)
}